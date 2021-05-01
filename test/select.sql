SELECT -- named parameter のような :コメント
	customer.customer_id,	first_name,	payment.*
FROM
	customer INNER JOIN payment ON payment.customer_id = customer.customer_id
WHERE /** ブロックコメントもOK */
  customer.first_name = "髙橋≦:𠮷野" 
	AND customer.last_name = :饅頭 OR payment.type IN (:予算)
	AND customer.age = :age AND payment.service = :service
ORDER BY payment_date;
