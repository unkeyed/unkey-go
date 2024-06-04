# V1MigrationsEnqueueKeysHash

Provide either `hash` or `plaintext`


## Fields

| Field                                                                                                  | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `Value`                                                                                                | *string*                                                                                               | :heavy_check_mark:                                                                                     | The hashed and encoded key                                                                             |
| `Variant`                                                                                              | [operations.V1MigrationsEnqueueKeysVariant](../../models/operations/v1migrationsenqueuekeysvariant.md) | :heavy_check_mark:                                                                                     | The algorithm for hashing and encoding, currently only sha256 and base64 are supported                 |