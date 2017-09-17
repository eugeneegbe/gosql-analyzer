# GoSQL-Analyzer

Go tool for analyzing SQL Queries to improve on Database Performance in DBMSs.



# Algorithm Abstract

We are starting with the MySQL DBMS at the start;

```psuedocode
console = Prepare_Analyzer_Console()

if console is ready

	Q = Get_SQL_Query()

	if Q is valid in MySQL
		R = Analyze_Query()
		return R

	else 
		Q is invalid in MySQL
		
	Enter a valid query
	goto line #14

else
	exit(console)
```



# Future features

* Add support for other DBMSs like Oracle, PostgreSQL etc.
