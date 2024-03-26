#include <stdio.h>
#include <stdlib.h>
#include "unreal_nibi_sdk.h"

int main()
{
    // Create a new NibiruClient instance using the exported Go function.
    int ret = NewNibiruClientDefault();
    if (ret != 0)
    {
        printf("Failed to create NibiruClient\n");
        return 1;
    }
    // Create new wallet
    // Generate Menomonic
    char *prases = GenerateRecoveryPhrase();
    printf("Prases %s", prases);

    // Create key(private,public =>signner) from menemonic
    int createAccount = CreateAccount(prases);
    if (createAccount != 0)
    {
        printf("Failed to create account\n");
        return 1;
    }
    // Storage in keyring
    int addSigner = AddSignerToKeyring(prases);
    if (addSigner != 0)
    {
        printf("Failed to add signer\n");
        return 1;
    }

    // Get wallet address
    int address = GetAddressFromMnemonic(prases);
    if (address != 0)
    {
        printf("Failed to get private key\n");
        return 1;
    }

    int privkey = GetPrivKeyFromMnemonic(prases);
    if (privkey != 0)
    {
        printf("Failed to get private key\n");
        return 1;
    }

    return 0;
}