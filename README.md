# logParser
Golang multiple log-files parser

The task: Implement a log converter to the required format.


There are several ways to determine the format of the Log Record. We can explicitly pass the format type along with the file or parse the lines, trying to isolate the correspondence to a specific format.
It is obvious that Log Formats differ only in the format of the Date.
To determine the Format of the Log Record, I decided to  iteratively parse the Date field for different views. This approach will allow us to correctly handle any type of Date fields and easily add other Log Formats.
All invalid entries (Lines) are ignored.

For each file, a separate thread (Goriutine) will be launched, which will check the state of the file at specific intervals. Checks the Last Modified fiels. For each file, the last read line number is stored, so only new data is transferred to the main hread.