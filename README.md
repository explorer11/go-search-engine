Basic text search engine in files.  
Written in Go.  
Given a directory, the search engine builds an index in memory.  
It looks for the words of a query in the content and in the title of the files in this folder and in all sub folders.  
The score of a file is the ratio of the words which are found in it : 1 if a file contains all the words, 0.5 if a file contains one in two words etc  
All files containing at least one word in a query are returned, with no order.  

To install, execute   
go install searchengine  

The program takes 2 arguments  
d : the directory containing the files  
m : the mode (default : http) : command or http  
When the mode is command, the program is a console application which takes a query in the command prompt  
When the mode is http, the program runs in a server on port 10000. A query can be made at this url : http://localhost:10000/search?q=myquery  

