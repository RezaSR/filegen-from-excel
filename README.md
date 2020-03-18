# filegen-from-excel
Generate text files from Excel data based on a template.

## Usage
Run the program with -h option to view usage:

```
-c	Run in CLI mode and do not open GUI
-d string
	Excel data file
-f string
	Output file name that contains special patterns:
	[0000]:
		Generates auto increment number padded to the specified zeros
		For example:
		[00].txt generates: 00.txt, 01.txt, 02.txt, 03.txt, ...
	[COLUMN]:
		Replaces with the content of corresponding column from excel data
		For example:
		[A].txt replaces [A] with the data of cell "A" of current row
	Patterns can be escaped by adding ":" after "["
		For example:
		[:00].txt generates [00].txt
	 (default "[0000].txt")
-o string
	Output directory (default "out")
-t string
	Template file that contains patterns to be replaced by excel data:
	[COLUMN]:
		Replaces with the content of corresponding column from excel data
		For example:
		[A] replaces with the data of cell "A" of current row
	Patterns can be escaped by adding ":" after "["
		For example:
		[:A] generates [A]
-v	Version number
```
