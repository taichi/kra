SELECT foo , bar FROM baz WHERE foo IN (SELECT kind FROM films WHERE kind IN (@予算, ?))
