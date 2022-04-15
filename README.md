# KPM2CSV
This tiny command-line tool helps you with moving from Kaspersky Password Manager to [KeePassXC](https://keepassxc.org/) by converting the exported credentials into a CSV format.

## How to get it?
The latest releases can be found [here](https://github.com/Th3Shadowbroker/KPM2CSV/releases) and include a compiled binary for Windows (amd64).

## How to use it?
1. Run `kpm2csv -i <exported file> -o <csv file>`
2. Open KeePassXC
3. Select `Database` > `Import` > `CSV-File`
4. Select the CSV file that was generated
5. Follow the instructions, until you reach the import settings
6. Select `TAB (\t)` as Separator and check the `First row contains field names` option
7. You can now find the imported credentials in the `Kaspersky Password Manager` group in KeePass

## Building from source
```
go mod download
go mod verify
go build .
```
