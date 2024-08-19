#include <termios.h>
#include "header.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sqlite3.h>

// Chemin au fichier de la database
char *USERS = "./data/users.txt";
static char *DB_FILE = "./data/users.sqlite";

// TODO : supp termios pour cahcer l'input user et utliser noecho() a la place
// Fonction qui gere le menu login
void loginMenu(struct User *u)
{
    //   Import qui permet de cacher
    struct termios oflags, nflags;

    system("clear");
    printf("\n\n\n\t\t\t\t   Bank Management System\n\t\t\t\t\t User Login:");
    scanf("%s", u->name);

    // Disabling echo, allowing to hide password
    tcgetattr(fileno(stdin), &oflags);
    nflags = oflags;
    nflags.c_lflag &= ~ECHO;
    nflags.c_lflag |= ECHONL;

    if (tcsetattr(fileno(stdin), TCSANOW, &nflags) != 0)
    {
        perror("tcsetattr");
        exit(1);
    }

    printf("\n\n\n\n\n\t\t\t\tEnter the password to login:");
    scanf("%s", u->password);

    // Appel la focntion qui verifie si le compte
    if (verifyCredentials(u) == 0)
    {
        printf("Wrong UserName or Password\n");
        exit(0);
    }
    printf("%d", u->id);

    //  Restore terminal
    if (tcsetattr(fileno(stdin), TCSANOW, &oflags) != 0)
    {
        perror("tcsetattr");
        exit(1);
    }
}

// Verifie si un compte est deja utilisé
int verifyCredentials(struct User *u)
{
    //  Open database connection
    sqlite3 *db;
    int rc = sqlite3_open(DB_FILE, &db);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "Cannot open or create the database: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    // Prepare SQL statement
    sqlite3_stmt *stmt;
    const char *query = "SELECT id, name, password FROM users WHERE name = ?";
    rc = sqlite3_prepare_v2(db, query, -1, &stmt, NULL);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "SQL error3: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    // Bind parameters
    sqlite3_bind_text(stmt, 1, u->name, -1, SQLITE_STATIC);
    sqlite3_bind_text(stmt, 2, u->password, -1, SQLITE_STATIC);

    // Execute SQL statement
    int count = 0;
    if (sqlite3_step(stmt) == SQLITE_ROW)
    {
        // Retrieve user ID and password from query result
        u->id = sqlite3_column_int(stmt, 0);
        const unsigned char *dbPassword = sqlite3_column_text(stmt, 2);
        if (strcmp((char *)dbPassword, u->password) == 0)
        {
            // User found and password matches
            count = 1;
        }
    }

    // Finalize statement and close database connection
    sqlite3_finalize(stmt);
    sqlite3_close(db);

    return count;
}

// Fonction qui gere le menu register
void registerMenu(struct User *u)
{
    //  Check if database file exists
    // Database file doesn't exist, create it
    sqlite3 *db;
    int rc = sqlite3_open(DB_FILE, &db);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "Cannot open or create the database: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    // Créer la table users
    const char *createTableSQL = "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, password TEXT);";
    rc = sqlite3_exec(db, createTableSQL, NULL, 0, NULL);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "Failed to create table: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    // Affcihage terminal
    system("clear");
    printf("\t\t====== Sign Up =====\n\n");
    printf("Enter User Name:");
    scanf("%49s", u->name);
    printf("Enter Password:");
    scanf("%49s", u->password);

    // Verifie si le nom est deja utilise
    /*     if (AlreadyExistAccount(u->name) == 1)
        {
            printf("Name and Password already used\n");
            sqlite3_close(db);
            return;
        } */
    if (verifyCredentials(u))
    {
        printf("Name and Password already used\n");
        sqlite3_close(db);
        exit(0);
    }

    // Insertion SQL
    const char *insertSQL = "INSERT INTO users (name, password) VALUES (?, ?);";
    sqlite3_stmt *stmt;
    rc = sqlite3_prepare_v2(db, insertSQL, -1, &stmt, NULL);

    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "SQL error1: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    sqlite3_bind_text(stmt, 1, u->name, -1, SQLITE_STATIC);
    sqlite3_bind_text(stmt, 2, u->password, -1, SQLITE_STATIC);

    rc = sqlite3_step(stmt);
    if (rc != SQLITE_DONE)
    {
        fprintf(stderr, "SQL error2: %s\n", sqlite3_errmsg(db));
        sqlite3_finalize(stmt);
        sqlite3_close(db);
        exit(1);
    }

    printf("User added successfully!\n");

    sqlite3_finalize(stmt);
    sqlite3_close(db);
}

// Verifie si le compte existe deja
int AlreadyExistAccount(struct User *u)
{
    sqlite3 *db;
    int rc = sqlite3_open(DB_FILE, &db);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "Cannot open or create the database: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    sqlite3_stmt *stmt;
    const char *query = "SELECT COUNT(*) FROM users WHERE name = ?; ";
    rc = sqlite3_prepare_v2(db, query, -1, &stmt, NULL);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "SQL error3: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        exit(1);
    }

    sqlite3_bind_text(stmt, 1, u->name, -1, SQLITE_STATIC);
    /* sqlite3_bind_text(stmt, 2, password, -1, SQLITE_STATIC); */

    int count = 0;
    if (sqlite3_step(stmt) == SQLITE_ROW)
    {
        count = sqlite3_column_int(stmt, 0);
    }

    sqlite3_finalize(stmt);
    sqlite3_close(db);

    return count > 0;
}
