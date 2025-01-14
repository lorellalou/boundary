// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"github.com/hashicorp/boundary/globals"
	"github.com/hashicorp/boundary/internal/credential"
	"github.com/hashicorp/boundary/internal/db"
	"github.com/hashicorp/boundary/internal/errors"
	"github.com/hashicorp/boundary/internal/types/subtypes"
)

func init() {
	if err := subtypes.Register(credential.Domain, Subtype, globals.VaultCredentialStorePrefix, DynamicCredentialPrefix); err != nil {
		panic(err)
	}
	if err := subtypes.Register(credential.Domain, GenericLibrarySubtype, globals.VaultCredentialLibraryPrefix); err != nil {
		panic(err)
	}
	if err := subtypes.Register(credential.Domain, SSHCertificateLibrarySubtype, globals.VaultSshCertificateCredentialLibraryPrefix); err != nil {
		panic(err)
	}
}

// PublicId prefixes for the resources in the vault package.
const (
	// DynamicCredentialPrefix is the prefix for Vault dynamic credentials
	DynamicCredentialPrefix = "cdvlt"

	Subtype                      = subtypes.Subtype("vault")
	GenericLibrarySubtype        = subtypes.Subtype("vault-generic")
	SSHCertificateLibrarySubtype = subtypes.Subtype("vault-ssh-certificate")
)

func newCredentialStoreId() (string, error) {
	id, err := db.NewPublicId(globals.VaultCredentialStorePrefix)
	if err != nil {
		return "", errors.WrapDeprecated(err, "vault.newCredentialStoreId")
	}
	return id, nil
}

func newCredentialId() (string, error) {
	id, err := db.NewPublicId(DynamicCredentialPrefix)
	if err != nil {
		return "", errors.WrapDeprecated(err, "vault.newCredentialId")
	}
	return id, nil
}

func newCredentialLibraryId() (string, error) {
	id, err := db.NewPublicId(globals.VaultCredentialLibraryPrefix)
	if err != nil {
		return "", errors.WrapDeprecated(err, "vault.newCredentialLibraryId")
	}
	return id, nil
}

func newSSHCertificateCredentialLibraryId() (string, error) {
	id, err := db.NewPublicId(globals.VaultSshCertificateCredentialLibraryPrefix)
	if err != nil {
		return "", errors.WrapDeprecated(err, "vault.newSSHCertificateCredentialLibraryPrefix")
	}
	return id, nil
}
