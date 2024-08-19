#include "header.h"
#include <sqlite3.h>
#include <ncurses.h>

const char *RECORDS = "./data/records.txt";
const char *DB_FILE = "./data/users.sqlite";

void stayOrReturn(int notGood, void f(struct User *u), struct User *u)
{
    int option;
    if (notGood == 0)
    {
        printf("\n✖ Record not found!!\n");
    invalid:
        printf("\nEnter 0 to try again, 1 to return to main menu and 2 to exit:");
        scanf("%d", &option);
        if (option == 0)
            f(u);
        else if (option == 1)
            mainMenu(&u);
        else if (option == 2)
            exit(0);
        else
        {
            printf("Insert a valid operation!\n");
            goto invalid;
        }
    }
    else
    {
        printf("\nEnter 1 to go to the main menu and 0 to exit:");
        scanf("%d", &option);
    }
    if (option == 1)
    {
        system("clear");
        mainMenu(&u);
    }
    else
    {
        system("clear");
        exit(1);
    }
}

// Fonction qui permet de passer directement au suivant si l'action est vrai
// TODO : Implanter ausssi
void success(struct User *u)
{
    int option;
    printf("\n✔ Success!\n\n");
invalid:
    printf("Enter 1 to go to the main menu and 0 to exit!\n");
    scanf("%d", &option);
    system("clear");
    if (option == 1)
    {
        mainMenu(u);
    }
    else if (option == 0)
    {
        exit(1);
    }
    else
    {
        printf("Insert a valid operation!\n");
        goto invalid;
    }
}

// Fonction qui creer un nouvel accompte
void createNewAcc(struct User *u)
{
    struct Record r;
    // ouverture de la database
    sqlite3 *db;
    int rc = sqlite3_open(DB_FILE, &db);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "Cannot open or create the database: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }
    // Création de la table records
    const char *createTableSQL = "CREATE TABLE IF NOT EXISTS records (id INTEGER PRIMARY KEY, userId INT,INTEGER, name TEXT, password TEXT,account INTEGER UNIQUE, depositDay INT, depositMonth INT, depositYear INT, country TEXT, phone INTEGER, balance REAL, accountType TEXT);";
    rc = sqlite3_exec(db, createTableSQL, NULL, 0, NULL);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "Failed to create table: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    system("clear");
    printf("\t\t\t===== New record =====\n");

    printf("\nEnter way's date(mm/dd/yyyy):");

    scanf("%d/%d/%d", &r.depositDay, &r.depositMonth, &r.depositYear);

    /*     if (sscanf(, "%d/%d/%d", &r.depositDay, &r.depositMonth, &r.depositYear))
        {
        } */
    printf("\nEnter the account number:");
    scanf("%d", &r.account);
    // Initiation de la requete
    sqlite3_stmt *stmt;
    // Cste qui permet de recupere le nom et le numero de compte
    const char *ask = "SELECT * FROM records WHERE name= ? AND account= ?";
    // execute la requete
    int req = sqlite3_prepare_v2(db, ask, -1, &stmt, NULL);
    // Lie les information au bon format et dans la bonne colonne
    sqlite3_bind_text(stmt, 1, u->name, -1, SQLITE_STATIC);
    sqlite3_bind_int(stmt, 2, r.account);
    // gestion erreur
    if (req != SQLITE_OK)
    {
        fprintf(stderr, "SQL error 4: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }
    // Permet de determiner si le compte existe deja
    if (sqlite3_step(stmt) == SQLITE_ROW)
    {
        printf("\n✖ This Account already exists for this user\n\n");
        sqlite3_finalize(stmt);
        sqlite3_close(db);
        exit(1);
        goto noAccount;
    }

// Pas de compte detecte
noAccount:
    printf("\nEnter the country:");
    scanf("%s", r.country);
    printf("\nEnter the phone number:");
    scanf("%d", &r.phone);
    printf("\nEnter amount to deposit: $");
    scanf("%lf", &r.balance);
    printf("\nChoose the type of account:\n\t-> saving\n\t-> current\n\t-> fixed01(for 1 year)\n\t-> fixed02(for 2 years)\n\t-> fixed03(for 3 years)\n\n\tEnter your choice:");
    scanf("%s", r.accountType);

    /*     // Convert string to double
        char *endptrBal;
        r.balance = strtod(balance_str, &endptrBal);
        if (*endptr != '\0')
        {
            // Input is not a valid double
            mvprintw(center_y + 2, center_x - 10, "Invalid input! Please enter a valid amount.");
            refresh();
            getch(); // Wait for user input
            // You might want to handle this case appropriately, e.g., ask for input again or exit the function
            return;
        } */

    // Envoie des données dans la table records
    const char *insertSQL = "INSERT INTO records (userId, name, password, account, depositDay, depositMonth, depositYear, country, phone, balance, accountType) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)";
    rc = sqlite3_prepare_v2(db, insertSQL, -1, &stmt, NULL);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "SQL error 1: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }
    // FIXME : id pas renvoyer
    sqlite3_bind_int(stmt, 1, u->id);
    sqlite3_bind_text(stmt, 2, u->name, -1, SQLITE_STATIC);
    sqlite3_bind_text(stmt, 3, u->password, -1, SQLITE_STATIC);
    sqlite3_bind_int(stmt, 4, r.account);
    sqlite3_bind_int(stmt, 5, r.depositDay);
    sqlite3_bind_int(stmt, 6, r.depositMonth);
    sqlite3_bind_int(stmt, 7, r.depositYear);
    sqlite3_bind_text(stmt, 8, r.country, -1, SQLITE_STATIC);
    sqlite3_bind_int(stmt, 9, r.phone);
    sqlite3_bind_double(stmt, 10, r.balance);
    sqlite3_bind_text(stmt, 11, r.accountType, -1, SQLITE_STATIC);

    rc = sqlite3_step(stmt);
    if (rc != SQLITE_DONE)
    {
        fprintf(stderr, "SQL error 2: %s\n", sqlite3_errmsg(db));
        sqlite3_finalize(stmt);
        sqlite3_close(db);
        exit(1);
    }

    sqlite3_finalize(stmt);
    sqlite3_close(db);

    success(u);
}

// Permet de voir tout les comptes
void checkAllAccounts(struct User *u)
{
    sqlite3 *db;
    int rc = sqlite3_open(DB_FILE, &db);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "Cannot open or create the database: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    // Selectionne dans la database
    sqlite3_stmt *stmt;
    const char *query = "SELECT account, depositDay , depositMonth , depositYear, country, phone, balance, accountType FROM records WHERE name=?;";
    int req = sqlite3_prepare_v2(db, query, -1, &stmt, NULL);

    if (req != SQLITE_OK)
    {
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    sqlite3_bind_text(stmt, 1, u->name, -1, SQLITE_STATIC);

    system("clear");
    printf("\t\t====== All accounts from user, %s =====\n\n", u->name);
    printf("_____________________\n");

    // Boucle qui permet de recupere les infos tant que le stmt prend une ligne
    while (sqlite3_step(stmt) == SQLITE_ROW)
    {
        struct Record r;
        r.account = sqlite3_column_int(stmt, 0);
        r.depositDay = sqlite3_column_int(stmt, 1);
        r.depositMonth = sqlite3_column_int(stmt, 2);
        r.depositYear = sqlite3_column_int(stmt, 3);
        const unsigned char *date = sqlite3_column_text(stmt, 4);
        const unsigned char *country = sqlite3_column_text(stmt, 5);
        r.phone = sqlite3_column_int(stmt, 6);
        r.balance = sqlite3_column_double(stmt, 7);
        const unsigned char *accountType = sqlite3_column_text(stmt, 7);

        printf("\nAccount number: %d\nDeposit Date: %d/%d/%d\nCountry: %s\nPhone number: %d\nAmount deposited: $%.2f\nType Of Account: %s\n",
               r.account,
               r.depositDay,
               r.depositMonth,
               r.depositYear,
               country,
               r.phone,
               r.balance,
               accountType);
    }

    sqlite3_finalize(stmt);
    sqlite3_close(db);
    success(u);
}

// Fonction menu qui permet de mettre a jour le numeros de telephone ou le pays
void UpdateAccount(struct User *u)
{
    struct Record r;                     // Structure record
    sqlite3 *db;                         // ouverture de la database
    int rc = sqlite3_open(DB_FILE, &db); // Prepare la declaration SQL
    if (rc != SQLITE_OK)
    {
        // Gestion d'erreur
        fprintf(stderr, "Cannot open or create the database: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    // initialise la variable option qui permet de choisir dans le menu
    int option;
    // Affichage
    printf("\nWhat is the account number you want to change: ");
    if (scanf("%d", &(r.account)) != 1)
    {
        // Getion d'erreur
        fprintf(stderr, "Invalid input for account number\n");
        sqlite3_close(db);
        exit(1);
    }

    // Affichage
    printf("\nWhich information do you want to update ?");
    printf("\n1-> Phone number");
    printf("\n2-> Country\n\n");
    if (scanf("%d", &option) != 1)
    {
        fprintf(stderr, "Invalid input for option\n");
        sqlite3_close(db);
        exit(1);
    }

    switch (option)
    {
    case 1:
        Updatephone(u, &r, db);
        break;
    case 2:
        UpdateCountry(u, &r, db);
        break;
    default:
        printf("Invalid operation!\n");
        sqlite3_close(db);
        exit(1);
    }
    sqlite3_close(db); // Feremeture de la database
    success(u);        // Appel de la fonction success
}

// Permet de mettre a jour seulement le numeros de telephone
void Updatephone(struct User *u, struct Record *r, sqlite3 *db)
{
    sqlite3_stmt *stmt;
    // Requete SQL qui permet de mettre a jour le telephone en fonction de l'utilisateur et du compte
    const char *repl = "UPDATE records SET phone=? WHERE name=? AND account=?; ";
    // Prepare la declaration SQL
    int req = sqlite3_prepare_v2(db, repl, -1, &stmt, NULL);
    if (req != SQLITE_OK)
    {
        // Gestion d'erreur
        fprintf(stderr, "SQL error0: %s\n", sqlite3_errmsg(db));
        sqlite3_finalize(stmt); // Termine la declaration SQL
        sqlite3_close(db);      // Ferme la database
        exit(1);                // Ferme le Programme avec une erreur
    }

    // Affcihage
    system("clear");
    printf("Enter the New phone number : ");
    if (scanf("%d", &(r->phone)) != 1)
    {
        // Gestion d'erreur
        fprintf(stderr, "Invalid input for phone number\n");
        sqlite3_finalize(stmt); // Termine la declaration SQL
        sqlite3_close(db);      // Ferme la database
        exit(1);                // Ferme le Programme avec une erreur
    }

    // Lie le nom, le nouveau numeros de telephone et le compte avec la declaration
    sqlite3_bind_int(stmt, 1, r->phone);
    sqlite3_bind_text(stmt, 2, u->name, -1, SQLITE_STATIC);
    sqlite3_bind_int(stmt, 3, r->account);

    // Execute la declaration
    req = sqlite3_step(stmt);
    if (req != SQLITE_DONE)
    {
        // Gestion d'erreur
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
    }
    // Finalise la declaration et ferme la databse
    sqlite3_finalize(stmt); // Termine la declaration SQL
    sqlite3_close(db);      // Ferme la database
}

// Fonction qui permet de mettre a jour le pays
void UpdateCountry(struct User *u, struct Record *r, sqlite3 *db)
{
    sqlite3_stmt *stmt;
    // Requete SQL qui permet de mettre a jour le pays en fonction de l'utilisateur et du compte
    const char *repl = "UPDATE records SET country=? WHERE name=? AND account=?;";
    // Prepare la declaration SQL
    int req = sqlite3_prepare_v2(db, repl, -1, &stmt, NULL);
    if (req != SQLITE_OK)
    {
        // Gestion d'erreur
        fprintf(stderr, "SQL error0: %s\n", sqlite3_errmsg(db));
        sqlite3_finalize(stmt); // Termine la declaration SQL
        sqlite3_close(db);      // Ferme la database
        exit(1);                // Ferme le Programme avec une erreur
    }
    // Affcihage
    system("clear");
    printf("Enter the New Country : ");
    scanf("%s", &r->country);

    // Lie le nom, le nouveau pays et le compte avec la declaration
    sqlite3_bind_text(stmt, 1, r->country, -1, SQLITE_STATIC);
    sqlite3_bind_text(stmt, 2, u->name, -1, SQLITE_STATIC);
    sqlite3_bind_int(stmt, 3, r->account);

    req = sqlite3_step(stmt);
    if (req != SQLITE_DONE)
    {
        // Gestion d'erreur
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
    }
    // Finalise la declaration et ferme la databse
    sqlite3_finalize(stmt); // Termine la declaration SQL
    sqlite3_close(db);      // Ferme la database
}

void CheckDetails(struct User *u)
{
    struct Record r;                     // Structure to hold record details
    sqlite3 *db;                         // Database connection
    int rc = sqlite3_open(DB_FILE, &db); // Open the database
    if (rc != SQLITE_OK)
    {
        // Error handling
        fprintf(stderr, "Cannot open or create the database: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    printf("\nEnter the account number: ");
    if (scanf("%d", &(r.account)) != 1)
    {
        // Error handling for invalid input
        fprintf(stderr, "Invalid input for account number\n");
        sqlite3_close(db);
        exit(1);
    }

    sqlite3_stmt *stmt;
    // SQL query to retrieve account details based on user and account number
    const char *query = "SELECT account, depositDay, depositMonth, depositYear, country, phone, balance, accountType FROM records WHERE name=? AND account=?;";
    // Prepare the SQL statement
    int req = sqlite3_prepare_v2(db, query, -1, &stmt, NULL);
    if (req != SQLITE_OK)
    {
        // Error handling for SQL preparation failure
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
        sqlite3_finalize(stmt); // Finalize the SQL statement
        sqlite3_close(db);      // Close the database
        exit(1);                // Exit the program with an error
    }

    system("clear");

    // Bind the account number to the SQL statement
    sqlite3_bind_text(stmt, 1, u->name, -1, SQLITE_STATIC);
    sqlite3_bind_int(stmt, 2, r.account);

    // Execute the SQL statement
    req = sqlite3_step(stmt);
    if (req != SQLITE_ROW)
    {
        // If no rows are returned, print a message and exit
        printf("Account details not found.\n");
        sqlite3_finalize(stmt); // Finalize the SQL statement
        sqlite3_close(db);      // Close the database
        success(u);             // Call success function with user
        return;
    }

    // Extract the details from the SQL result
    r.account = sqlite3_column_int(stmt, 0);
    r.depositDay = sqlite3_column_int(stmt, 1);
    r.depositMonth = sqlite3_column_int(stmt, 2);
    r.depositYear = sqlite3_column_int(stmt, 3);
    const unsigned char *country = sqlite3_column_text(stmt, 4);
    r.phone = sqlite3_column_int(stmt, 5);
    r.balance = sqlite3_column_double(stmt, 6);
    const unsigned char *accountType = sqlite3_column_text(stmt, 7);

    // Print the details
    printf("\nAccount number: %d\nDeposit Date: %d/%d/%d\nCountry: %s\nPhone number: %d\nAmount deposited: $%.2f\nType Of Account: %s\n",
           r.account,
           r.depositDay,
           r.depositMonth,
           r.depositYear,
           country,
           r.phone,
           r.balance,
           accountType);

    // Permet de savoir quel type d'investissement a le compte
    const char *invest = (const char *)sqlite3_column_text(stmt, 7);
    // Compare si l'investissement est savings
    if (strcmp(invest, "saving") == 0)
    {
        // Calcul l'interet
        double interest = ((r.balance / 100) * 7) / 12;
        fprintf(stdout, "\nYou will get $%f as interest on day %d of every month !\n", interest, r.depositDay);
    }
    // Compare si l'investissement est fixed01
    else if (strcmp(invest, "fixed01") == 0)
    {
        // Calcul l'interet
        double interest = (r.balance / 100) * 4;
        fprintf(stdout, "\nYou will get $%f as interest on %d/%d/%d  !\n", interest, r.depositDay, r.depositMonth, r.depositYear + 1);
    }
    // Compare si l'investissement est fixed02
    else if (strcmp(invest, "fixed02") == 0)
    {
        // Calcul l'interet
        double interest = ((r.balance / 100) * 5) * 2;
        fprintf(stdout, "\nYou will get $%f as interest on %d/%d/%d  !\n", interest, r.depositDay, r.depositMonth, r.depositYear + 2);
    }
    // Compare si l'investissement est fixed03
    else if (strcmp(invest, "fixed03") == 0)
    {
        // Calcul l'interet
        double interest = ((r.balance / 100) * 8) * 3;
        fprintf(stdout, "\nYou will get $%f as interest on %d/%d/%d  !\n", interest, r.depositDay, r.depositMonth, r.depositYear + 3);
    }
    // Compare si l'investissement est current
    else if (strcmp(invest, "current") == 0)
    {
        fprintf(stdout, "\nYou will not get interests because the account is of type current\n");
    }

    // Finalize the SQL statement
    sqlite3_finalize(stmt);
    // Close the database
    sqlite3_close(db);
    // Call success function with user
    success(u);
}

void MakeTransaction(struct User *u)
{
    struct Record r;                     // Structure record
    sqlite3 *db;                         // ouverture de la database
    int rc = sqlite3_open(DB_FILE, &db); // Prepare la declaration SQL
    if (rc != SQLITE_OK)
    {
        // Gestion d'erreur
        fprintf(stderr, "Cannot open or create the database: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }
    sqlite3_stmt *stmt;
    // SQL query to retrieve account details based on user and account number
    const char *query = "SELECT accountType, balance FROM records WHERE name=? AND account=?;";
    // Prepare the SQL statement
    int req = sqlite3_prepare_v2(db, query, -1, &stmt, NULL);
    if (req != SQLITE_OK)
    {
        // Error handling for SQL preparation failure
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
        sqlite3_finalize(stmt); // Finalize the SQL statement
        sqlite3_close(db);      // Close the database
        exit(1);                // Exit the program with an error
    }

    // Affichage
    printf("\nEnter the account number: ");
    if (scanf("%d", &(r.account)) != 1)
    {
        // Getion d'erreur
        fprintf(stderr, "Invalid input for account number\n");
        sqlite3_close(db);
        exit(1);
    }
    sqlite3_bind_text(stmt, 1, u->name, -1, SQLITE_STATIC);
    sqlite3_bind_int(stmt, 2, r.account);

    req = sqlite3_step(stmt);
    if (req != SQLITE_ROW)
    {
        fprintf(stderr, "Account details not found.\n");
        sqlite3_finalize(stmt);
        sqlite3_close(db);
        stayOrReturn(0, MakeTransaction, u);
    }

    const char *notransaction = (const char *)sqlite3_column_text(stmt, 0);
    r.balance = sqlite3_column_double(stmt, 1);

    if (strcmp(notransaction, "fixed01") == 0 ||
        strcmp(notransaction, "fixed02") == 0 ||
        strcmp(notransaction, "fixed03") == 0)
    {
        fprintf(stdout, "\nYour account type '%s' does not allow transactions.\n", notransaction);
        sqlite3_close(db);
        stayOrReturn(0, MakeTransaction, u);
        // exit(1);
    }

    // initialise la variable option qui permet de choisir dans le menu
    int option;
    // Affichage
    printf("\nDo you want to :");
    printf("\n1-> Withdraw");
    printf("\n2-> Deposit\n\n");
    printf("Enter you choice : ");
    if (scanf("%d", &option) != 1)
    {
        fprintf(stderr, "Invalid input for option\n");
        sqlite3_close(db);
        exit(1);
    }

    switch (option)
    {
    case 1:
        Withdraw(u, &r, db);
        break;
    case 2:
        Deposit(u, &r, db);
        break;
    default:
        printf("Invalid operation!\n");
        sqlite3_close(db);
        exit(1);
    }
    sqlite3_finalize(stmt);
    sqlite3_close(db); // Feremeture de la database
    success(u);        // Appel de la fonction success
}

void Withdraw(struct User *u, struct Record *r, sqlite3 *db)
{
    sqlite3_stmt *stmt;
    const char *repl = "UPDATE records SET balance = balance - ? WHERE name=? AND account=?;";
    int req = sqlite3_prepare_v2(db, repl, -1, &stmt, NULL);
    if (req != SQLITE_OK)
    {
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
        sqlite3_finalize(stmt);
        sqlite3_close(db);
        exit(1);
    }

    system("clear");
    printf("Enter the amount you want to withdraw: $");
    double withdrawAmount;
    if (scanf("%lf", &withdrawAmount) != 1)
    {
        fprintf(stderr, "Invalid input for withdrawal amount\n");
        sqlite3_finalize(stmt);
        sqlite3_close(db);
        exit(1);
    }

    if (r->balance < withdrawAmount)
    {
        fprintf(stderr, "\n✖ The amount you choose to withdraw is superior to your aviable balance!\n");
        sqlite3_finalize(stmt);
        sqlite3_close(db);
        exit(1);
    }

    sqlite3_bind_double(stmt, 1, withdrawAmount);
    sqlite3_bind_text(stmt, 2, u->name, -1, SQLITE_STATIC);
    sqlite3_bind_int(stmt, 3, r->account);

    req = sqlite3_step(stmt);
    if (req != SQLITE_DONE)
    {
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
    }

    sqlite3_finalize(stmt);
    sqlite3_close(db);
}

void Deposit(struct User *u, struct Record *r, sqlite3 *db)
{
    sqlite3_stmt *stmt;
    const char *repl = "UPDATE records SET balance = balance + ? WHERE name=? AND account=?;";
    int req = sqlite3_prepare_v2(db, repl, -1, &stmt, NULL);
    if (req != SQLITE_OK)
    {
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
        sqlite3_finalize(stmt);
        sqlite3_close(db);
        exit(1);
    }

    system("clear");
    printf("Enter the amount you want to deposit: $");
    double depositAmount;
    if (scanf("%lf", &depositAmount) != 1)
    {
        fprintf(stderr, "Invalid input for depsosit amount\n");
        sqlite3_finalize(stmt);
        sqlite3_close(db);
        exit(1);
    }
    if (depositAmount < 0)
    {
        fprintf(stderr, "Invalid input for depsosit amount\n");
        sqlite3_finalize(stmt);
        sqlite3_close(db);
        exit(1);
    }

    sqlite3_bind_double(stmt, 1, depositAmount);
    sqlite3_bind_text(stmt, 2, u->name, -1, SQLITE_STATIC);
    sqlite3_bind_int(stmt, 3, r->account);

    req = sqlite3_step(stmt);
    if (req != SQLITE_DONE)
    {
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
    }

    sqlite3_finalize(stmt);
    sqlite3_close(db);
}

void RemoveAccount(struct User *u)
{
    struct Record r;
    sqlite3 *db;
    sqlite3_stmt *stmt = NULL;
    int rc;

    rc = sqlite3_open(DB_FILE, &db);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "Cannot open or create the database: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    printf("\nEnter the account number you want to delete: ");
    if (scanf("%d", &(r.account)) != 1)
    {
        fprintf(stderr, "Invalid input for account number\n");
        sqlite3_close(db);
        exit(1);
    }

    // Check if the account exists and belongs to the user
    const char *checkQuery = "SELECT `account` FROM `records` WHERE `name`=? AND `account`=?";
    rc = sqlite3_prepare_v2(db, checkQuery, -1, &stmt, NULL);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    sqlite3_bind_text(stmt, 1, u->name, -1, SQLITE_STATIC);
    sqlite3_bind_int(stmt, 2, r.account);

    rc = sqlite3_step(stmt);
    if (rc != SQLITE_ROW)
    {
        printf("Account not found or does not belong to the user.\n");
        sqlite3_finalize(stmt);
        sqlite3_close(db);
        return;
    }

    sqlite3_finalize(stmt);

    // Proceed with account deletion
    const char *deleteQuery = "DELETE FROM `records` WHERE `name`=? AND `account`=?";
    rc = sqlite3_prepare_v2(db, deleteQuery, -1, &stmt, NULL);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    sqlite3_bind_text(stmt, 1, u->name, -1, SQLITE_STATIC);
    sqlite3_bind_int(stmt, 2, r.account);

    rc = sqlite3_step(stmt);
    if (rc == SQLITE_DONE)
    {
        printf("Success deleting account.\n");
    }
    else
    {
        fprintf(stderr, "Failed to delete account: %s\n", sqlite3_errmsg(db));
    }

    sqlite3_finalize(stmt);
    sqlite3_close(db);
    success(u);
}

void Transferowner(struct User *u)
{
    struct Record r;
    sqlite3 *db;
    sqlite3_stmt *stmt = NULL;
    sqlite3_stmt *stmt2 = NULL;
    sqlite3_stmt *stmt3 = NULL;
    int rc;

    rc = sqlite3_open(DB_FILE, &db);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "Cannot open or create the database: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    printf("\nEnter the account number you want to transfer ownership: ");
    if (scanf("%d", &(r.account)) != 1)
    {
        fprintf(stderr, "Invalid input for account number\n");
        sqlite3_close(db);
        exit(1);
    }

    const char *query = "SELECT `account`, `depositDay`, `depositMonth`, `depositYear`, `country`, `phone`, `balance`, `accountType`, `Id` FROM `records` WHERE `account`=?";
    rc = sqlite3_prepare_v2(db, query, -1, &stmt2, NULL);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    sqlite3_bind_int(stmt2, 1, r.account);
    rc = sqlite3_step(stmt2);
    if (rc != SQLITE_ROW)
    {
        printf("Account details not found.\n");
        sqlite3_finalize(stmt2);
        sqlite3_close(db);
        return;
    }

    r.account = sqlite3_column_int(stmt2, 0);
    r.depositDay = sqlite3_column_int(stmt2, 1);
    r.depositMonth = sqlite3_column_int(stmt2, 2);
    r.depositYear = sqlite3_column_int(stmt2, 3);
    strncpy(r.country, (const char *)sqlite3_column_text(stmt2, 4), sizeof(r.country) - 1);
    r.country[sizeof(r.country) - 1] = '\0'; // Ensure null-termination
    r.phone = sqlite3_column_int(stmt2, 5);
    r.balance = sqlite3_column_double(stmt2, 6);
    strncpy(r.accountType, (const char *)sqlite3_column_text(stmt2, 7), sizeof(r.accountType) - 1);
    r.accountType[sizeof(r.accountType) - 1] = '\0'; // Ensure null-termination

    printf("\nAccount number: %d\nDeposit Date: %d/%d/%d\nCountry: %s\nPhone number: %d\nAmount deposited: $%.2f\nType Of Account: %s\n",
           r.account,
           r.depositDay,
           r.depositMonth,
           r.depositYear,
           r.country,
           r.phone,
           r.balance,
           r.accountType);

    sqlite3_finalize(stmt2);

    char newuser[100]; // Adjust size as needed
    printf("\nEnter the username of the new owner: ");
    if (scanf("%99s", newuser) != 1)
    {
        fprintf(stderr, "Invalid input for new owner's username\n");
        sqlite3_close(db);
        exit(1);
    }

    // Check if the new owner exists and get their user ID
    const char *userCheckQuery = "SELECT `id` FROM `users` WHERE `name`=?";
    rc = sqlite3_prepare_v2(db, userCheckQuery, -1, &stmt3, NULL);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    sqlite3_bind_text(stmt3, 1, newuser, -1, SQLITE_STATIC);
    rc = sqlite3_step(stmt3);
    if (rc != SQLITE_ROW)
    {
        printf("New owner does not exist.\n");
        sqlite3_finalize(stmt3);
        sqlite3_close(db);
        return;
    }

    int newUserId = sqlite3_column_int(stmt3, 0);
    sqlite3_finalize(stmt3);

    // Update the record with the new owner's name and user ID
    const char *updateQuery = "UPDATE `records` SET `name`=?, `userId`=? WHERE `account`=?";
    rc = sqlite3_prepare_v2(db, updateQuery, -1, &stmt, NULL);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    sqlite3_bind_text(stmt, 1, newuser, -1, SQLITE_STATIC);
    sqlite3_bind_int(stmt, 2, newUserId);
    sqlite3_bind_int(stmt, 3, r.account);

    rc = sqlite3_step(stmt);
    if (rc == SQLITE_DONE)
    {
        printf("Ownership transferred successfully.\n");
    }
    else
    {
        fprintf(stderr, "Failed to transfer ownership: %s\n", sqlite3_errmsg(db));
    }

    sqlite3_finalize(stmt);
    sqlite3_close(db);
    success(u);
}