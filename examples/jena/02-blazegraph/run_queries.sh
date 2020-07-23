#!/usr/bin/env bash

run_query='../common/run_query.sh'
data_file='../data/address-book.jsonld'

bash ${run_query} ${data_file} Q1 "List all triples in blazegraph." << END_QUERY
PREFIX ab: <http://learningsparql.com/ns/addressbook#> 
CONSTRUCT 
{ ?s ?p ?o }
FROM <http://172.17.0.1:9999/blazegraph/sparql>
WHERE 
{ ?s ?p ?o }
END_QUERY

bash ${run_query} ${data_file} Q2 "Select all triples in blazegraph." << END_QUERY
PREFIX ab: <http://learningsparql.com/ns/addressbook#> 
SELECT ?s ?p ?o 
WHERE
{
    SERVICE <http://172.17.0.2:9999/blazegraph/sparql> 
    {
        SELECT ?s ?p ?o 
        WHERE {}
    }
}
END_QUERY