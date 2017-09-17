# Go SQL Query Analyzer

The tool for analyzing SQL Queries to improve on Database Performance in DBMSs written in Go.



# Algorithm Abstract

We are starting with the MySQL DBMS at the start;

```psuedocode
console = Prepare_Analyzer_Console()

if console is ready

	start: Q = Get_SQL_Query()

	if Q is valid in MySQL
		R = Analyze_Query()
		return R

	else 
		Q is invalid in MySQL
		
	Enter a valid query
	goto start

else
	exit(console)
```



# Future features

* Add support for other DBMSs like Oracle, PostgreSQL etc.
