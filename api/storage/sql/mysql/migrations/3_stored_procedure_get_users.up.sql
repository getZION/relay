CREATE PROCEDURE get_users()
BEGIN

SELECT 
	JSON_ARRAYAGG(JSON_OBJECT(
		'id', u.id, 
        'did', u.did, 
        'username', u.username,
        'email', u.email,
        'name', u.name,
        'bio', u.bio,
        'img', u.img,
        'priceToMessage', u.price_to_message,
        'created', u.created,
        'updated', u.updated)
	) as 'Result'
FROM 
	users u;

END