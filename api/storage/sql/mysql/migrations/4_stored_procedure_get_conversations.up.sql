CREATE PROCEDURE get_conversations()
BEGIN

SELECT 
	JSON_ARRAYAGG(JSON_OBJECT(
		'id', c.id, 
        'zid', c.zid,
        'communityZid', c.community_zid,
        'userDid', c.user_did,
        'text', c.text,
        'link', c.link,
        'img', c.img,
        'video', c.video,
        'public', c.public,
        'publicPrice', c.public_price,
        'created', c.created,
        'updated', c.updated,
        'deleted', c.deleted,
        'comments', co.comments)
	) as 'Result'
FROM 
	conversations c
LEFT JOIN (
	SELECT 
		conversation_zid, 
		JSON_ARRAYAGG(JSON_OBJECT(
			'id', co.id, 
            'zid', co.zid,
            'userDid', co.user_did,
            'text', co.text,
            'link', co.link,
            'created', co.created,
            'updated', co.updated,
            'deleted', co.deleted)
		) AS comments
	FROM comments co
	GROUP BY co.conversation_zid
) co ON co.conversation_zid = c.zid;

END