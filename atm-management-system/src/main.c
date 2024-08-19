#include "header.h"
#include <stdlib.h>
#include <signal.h>

struct User *u;

void mainMenu(struct User *u)
{
    int option;

    system("clear");
    printf("\n\n\t\t======= ATM =======\n\n");
    printf("\n\t\t-->> Feel free to choose one of the options below <<--\n");
    printf("\n\t\t[1]- Create a new account\n");
    printf("\n\t\t[2]- Update account information\n");
    printf("\n\t\t[3]- Check accounts\n");
    printf("\n\t\t[4]- Check list of owned account\n");
    printf("\n\t\t[5]- Make Transaction\n");
    printf("\n\t\t[6]- Remove existing account\n");
    printf("\n\t\t[7]- Transfer ownership\n");
    printf("\n\t\t[8]- Exit\n");
    scanf("%d", &option);

    switch (option)
    {
    case 1:
        createNewAcc(u);
        break;
    case 2:
        UpdateAccount(u);
        break;
    case 3:
        CheckDetails(u);
        break;
    case 4:
        checkAllAccounts(u);
        break;
    case 5:
        MakeTransaction(u);
        break;
    case 6:
        RemoveAccount(u);
        break;
    case 7:
        Transferowner(u);
        break;
    case 8:
        exit(1);
        break;
    default:
        printf("Invalid operation!\n");
        break;
    }
}
void initMenu(struct User *u)
{
    int r = 0;
    int option;
    system("clear");
    printf("\n\n\t\t======= ATM =======\n");
    printf("\n\t\t-->> Feel free to login / register :\n");
    printf("\n\t\t[1]- login\n");
    printf("\n\t\t[2]- register\n");
    printf("\n\t\t[3]- exit\n");
    while (!r)
    {
        scanf("%d", &option);
        printf("%d", option);
        switch (option)
        {
        case 1:
        {
            loginMenu(u);
            printf("%s %s", u->name, u->password);
            r = 1;
            break;
        }
        case 2:
            registerMenu(u);
            r = 1;
            break;
        case 3:
            exit(1);
            break;
        default:
            printf("Insert a valid operation!\n");
        }
    }
}

int main()
{
    struct User u;

    initMenu(&u);
    mainMenu(&u);
    return 0;
}
