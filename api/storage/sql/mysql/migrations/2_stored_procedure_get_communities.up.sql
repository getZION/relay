CREATE PROCEDURE get_communities()
BEGIN

SELECT 
	JSON_ARRAYAGG(JSON_OBJECT(
		'id', c.id, 
        'name', c.name, 
        'zid', c.zid, 
        'ownerDid', c.owner_did, 
        'ownerUsername', c.owner_username, 
        'description', c.description, 
        'escrowAmount', c.escrow_amount, 
        'img', c.img, 
        'lastactive', c.last_active, 
        'pricePerMessage', c.price_per_message, 
        'priceToJoin', c.price_to_join, 
        'public', c.public, 
        'created', c.created, 
		'updated', c.updated, 
        'deleted', c.deleted, 
        'tags', t.tags, 
        'users', u.users)
	) as 'Result'
FROM 
	communities c
LEFT JOIN (
	SELECT 
		community_zid, 
		JSON_ARRAYAGG(ct.tag) AS tags
	FROM community_tags ct
	GROUP BY ct.community_zid
) t ON t.community_zid = c.zid
LEFT JOIN (
	SELECT 
	community_zid, 
	JSON_ARRAYAGG(JSON_OBJECT(
		'userDid', cu.user_did,
        'joinedDate', cu.joined_date,
        'leftDate', cu.left_date,
        'leftReason', cu.left_reason)
	) AS users
    FROM community_users cu
    GROUP BY cu.community_zid
) u ON u.community_zid = c.zid;

END
