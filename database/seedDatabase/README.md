# Database Seeding

The `database/seedDatabase/content/` directory contains all the json files of the tables to be seeded. The file name should be of the format `table_name.json`

## How to Create Seed for a New Table:

1. Create a new file `table_name.json` in `database/seedDatabase/content/` directory.

2. Add a JSON object similar to the other seeder JSON files.

3. The JSON object should be of this type.
```json
{
    "table_name": [
        {
            "id": 1
        },
        {
            "id": 2
        }
    ]
}
```

4. `id` denotes the primary_key of the table.

5. Each new object in the array of `table_name` would represent a new row with `key` and `value` pairs as `coulmn_name` and `value` respectively.

6. After creating the file, go to `database/seedDatabase/seed.go`, add that `table_name` in the `seeds` or `devSeeds` variable (wherever appropriate) in the `main()` function.