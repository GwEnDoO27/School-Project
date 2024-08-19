#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sqlite3.h>

// all fields for each record of an account
struct Record
{
    int id;
    int userId;
    char name[100];
    char country[100];
    int phone;
    char accountType[10];
    int account;
    double balance;
    int depositDay;
    int depositMonth;
    int depositYear;
    char withdraw;
};

struct User
{
    int id;
    char name[50];
    char password[50];
};

// authentication functions
void loginMenu(struct User *u);
int verifyCredentials(struct User *u);
int AlreadyExistAccount(struct User *u);
void registerMenu(struct User *u);
// const char *getPassword(struct User u);

// system function
void createNewAcc(struct User *u);
void mainMenu(struct User *u);
void checkAllAccounts(struct User *u);
void UpdateAccount(struct User *u);
void UpdateCountry(struct User *u, struct Record *r, sqlite3 *db);
void Updatephone(struct User *u, struct Record *r, sqlite3 *db);
void CheckDetails(struct User *u);
void MakeTransaction(struct User *u);
void Deposit(struct User *u, struct Record *r, sqlite3 *db);
void Withdraw(struct User *u, struct Record *r, sqlite3 *db);
void RemoveAccount(struct User *u);
void Transferowner(struct User *u);
void initMenu(struct User *u);