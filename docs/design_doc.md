# Go SQL Query Analyzer

The tool for analyzing SQL Queries to improve on Database Performance in DBMSs written in Go.



# Algorithm Abstract

We are starting with the MySQL DBMS at the start;

```psuedocode
console = Prepare_Analyzer_Console()

if console is ready

	start: Q = Get_SQL_Query()

	if Q is valid in MySQL
		M = Analyze_Query()
		
		if Q can be optimized
			OQ = Optimize_Query()
			return OQ

		return M, Q

	else 
		Q is invalid in MySQL
		
	Enter a valid query
	goto start

else
	exit(console)
```

Meaning of the following letters in the pseudocode;
* M = Metrics of the Q (query)
* Q = Query
* OQ = Optimized Query



# Analyse_Query() method

This section analyzes the query and returns a bunch of metrics useful to determine how the query can be made efficient such as;

* Query Analyzer Table
* Query Response Time Index(QRTi)
* Response Statistics
* Example Query
* Explain Query
* Graphs

For more information, check the [MySQL Enterprise Monitor](https://www.mysql.com/products/enterprise/query.html).



# Future features

* Add support for other DBMSs like Oracle, PostgreSQL etc.
