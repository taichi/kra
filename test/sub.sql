SELECT foo , bar FROM baz WHERE foo IN (SELECT kind FROM films WHERE kind = 'CDR' OR kind = 'ZDE')
