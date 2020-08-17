SELECT
 e.Email
FROM
(SELECT 
    Email, 
    COUNT(Email)
FROM
    Person
GROUP BY Email
HAVING COUNT(Email) > 1) AS e
