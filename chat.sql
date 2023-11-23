CREATE TABLE `users` (
  `id` varchar(100) NOT NULL,
  `username` varchar(20) DEFAULT NULL,
  `email` varchar(100) NOT NULL,
  `password` longtext NOT NULL,
  `description` varchar(50) DEFAULT NULL,
  `image` varchar(20) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email_unique` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci 




CREATE TABLE `contacts` (
  `id` varchar(100) NOT NULL,
  `user_id` varchar(100) NOT NULL,
  `username` varchar(20) DEFAULT NULL,
  `email` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_contacts_users` (`user_id`),
  CONSTRAINT `fk_contacts_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

-- example test insert data
insert into contacts (contact_id, user_id, username, email) 
values("test-1", "bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9", "firjissie", "firji@irawan");





CREATE TABLE `chats` (
  `id` varchar(50) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `chat_type` enum('group','1 on 1') NOT NULL DEFAULT '1 on 1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci 
-- exampl insert into chats
insert into chats(chat_id)
values("chat-1")




CREATE TABLE `messages` (
  `id` varchar(50) NOT NULL,
  `chat_id` varchar(50) NOT NULL,
  `sender_id` varchar(100) NOT NULL,
  `content` longtext,
  `status` enum('belum terkirim','terkirim','dibaca') NOT NULL DEFAULT 'belum terkirim',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  CONSTRAINT `fk_messages_chats` FOREIGN KEY (`chat_id`) REFERENCES `chats` (`id`),
  CONSTRAINT `fk_messages_users` FOREIGN KEY (`sender_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci 
-- example insert into messages
insert into messages(message_id, chat_id, sender_id)
values("message-1","chat-1", "bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9")


CREATE TABLE `message_recipients` (
  `id` varchar(50) NOT NULL,
  `message_id` varchar(50) NOT NULL,
  `recipient_id` varchar(50) NOT NULL,
  `status` enum('belum terkirim','terkirim','dibaca') NOT NULL DEFAULT 'belum terkirim',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_message_recipients_messages` FOREIGN KEY (`message_id`) REFERENCES `messages` (`id`),
  CONSTRAINT `fk_message_recipients_users` FOREIGN KEY (`recipient_id`) REFERENCES `users` (`id`)
)



ALTER TABLE contacts
       ADD CONSTRAINT `fk_contacts_users`
            FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)

DELIMITER //
CREATE TRIGGER before_insert_trigger
BEFORE INSERT ON contacts
FOR EACH ROW
BEGIN
  -- Mengonversi nilai AUTO_INCREMENT dengan prefix
  SET NEW.nilai = CONCAT('USD-', NEW.id);
END;
//
DELIMITER ;
CREATE TABLE notif (
	notif_id string not null
)

alter table chats

-- untuk detail chat
SELECT 
	c.id AS chat_id, 
	c.chat_type AS chat_type,  
	m.sender_id AS sender_id, 
	mr.recipient_id AS recipient_id, 
	m.content AS last_message, 
	m.status as status_pesan_pengirim, 
	mr.status as status_pesan_penerima, 
	m.created_at as message_created,
	mr.created_at AS message_recipient_created
FROM message_recipients  as mr 
JOIN messages m ON (m.id = mr.message_id)
JOIN chats c ON (c.id = m.chat_id)
WHERE c.id = '72b84e9c-f65e-4835-aa58-2a1f2185fff4'
ORDER BY 
	CASE 
		WHEN m.updated_at is not NULL THEN m.updated_at 
		ELSE m.created_at 
	END
DESC 

SELECT c.id AS chat_id, c.chat_type AS chat_type, m.sender_id AS sender_id, mr.recipient_id AS recipient_id, m.content AS last_message, m.status as status_pesan_pengirim, mr.status as status_pesan_penerima, m.created_at AS message_created, mr.created_at AS message_recipient_created
FROM message_recipients AS mr 
JOIN messages m ON (mr.message_id = m.id)
JOIN chats c ON (c.id = m.chat_id)
WHERE (mr.recipient_id = "50766a0c-d5b0-4e0f-9cb9-489177f222ef" OR mr.recipient_id = "bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9") 
    AND (m.sender_id = "50766a0c-d5b0-4e0f-9cb9-489177f222ef" OR m.sender_id = "bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9")
ORDER BY mr.created_at, m.created_at ASC 

-- untuk list chat
SELECT 
	c.id as chat_id, 
	c.chat_type as 'type', 
	m.sender_id as sender_id, 
	mr.recipient_id as recipient_id, 
	m.content as last_message, 
	m.status as "status_content_send", 
	mr.status as 'status_send_recieve', 
	m.created_at as 'created_at'
FROM message_recipients mr 
JOIN messages m ON (mr.message_id = m.id)
JOIN chats c ON (c.id = m.chat_id)
WHERE (m.sender_id = "bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9" OR mr.recipient_id = "bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9")
GROUP BY c.id , m.sender_id , mr.recipient_id, m.status, m.content, mr.status, m.created_at
ORDER BY MAX(m.created_at) DESC;

-- untuk list chat new
SELECT 
	c.id as chat_id,
	(SELECT mm.sender_id  
		FROM message_recipients mrr
		JOIN messages mm ON (mrr.message_id = mm.id) 
		WHERE mrr.id = MAX(mr.id)) as sender_id,
	(SELECT mrr.recipient_id 
		FROM message_recipients mrr
		JOIN messages mm ON (mrr.message_id = mm.id) 
		WHERE mrr.id = MAX(mr.id)) as recipient_id,
	(SELECT mm.content 
		FROM message_recipients mrr
		JOIN messages mm ON (mrr.message_id = mm.id) 
		WHERE mrr.id = MAX(mr.id)) as last_message,
	(SELECT mm.created_at  
		FROM message_recipients mrr
		JOIN messages mm ON (mrr.message_id = mm.id) 
		WHERE mrr.id = MAX(mr.id)) as content_send,
	CASE
        WHEN (SELECT mrr.recipient_id 
        	FROM message_recipients mrr 
        	JOIN messages mm ON (mrr.message_id = mm.id) 
        	WHERE mrr.id = MAX(mr.id)) = "bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9" 
        		THEN (
        				SELECT username 
        					FROM users WHERE id = (
        					SELECT mm.sender_id  
								FROM message_recipients mrr
								JOIN messages mm ON (mrr.message_id = mm.id) 
								WHERE mrr.id = MAX(mr.id)
				))
        ELSE (
        				SELECT username 
        					FROM users WHERE id = (SELECT mrr.recipient_id 
		FROM message_recipients mrr
		JOIN messages mm ON (mrr.message_id = mm.id) 
		WHERE mrr.id = MAX(mr.id)))
    END AS username
FROM message_recipients mr 
JOIN messages m ON (mr.message_id = m.id)
JOIN chats c ON (c.id = m.chat_id)
WHERE (m.sender_id = "bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9" OR mr.recipient_id = "bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9")
GROUP BY c.id
ORDER BY content_send DESC 

-- list chats version 2
SELECT
    c.id AS chat_id,
    mm.sender_id,
    mrr.recipient_id,
    mm.content AS last_message,
    mm.created_at AS content_send,
    CASE
        WHEN mrr.recipient_id = 'bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9' THEN (SELECT username FROM users WHERE id = mm.sender_id)
        ELSE (SELECT username FROM users WHERE id = mrr.recipient_id)
    END AS username
FROM (
    SELECT
        m.chat_id,
        MAX(m.id) AS max_message_id
    FROM message_recipients mr
    JOIN messages m ON mr.message_id = m.id
    WHERE (m.sender_id = 'bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9' OR mr.recipient_id = 'bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9')
    GROUP BY m.chat_id
) max_message
JOIN message_recipients mrr ON max_message.max_message_id = mrr.message_id
JOIN messages mm ON max_message.max_message_id = mm.id
JOIN chats c ON c.id = max_message.chat_id
ORDER BY content_send DESC;

-- list chats version 3
SELECT 
	chat.id AS chat_id,
	m2.sender_id AS sender_id,
	m2.content AS last_message,
	mr2.recipient_id  AS recipient_id,
	CASE
        WHEN mr2.recipient_id = 'bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9' THEN (SELECT username FROM users WHERE id = m2.sender_id)
        ELSE (SELECT username FROM users WHERE id = mr2.recipient_id)
    END AS username,
	m2.created_at
FROM (
	SELECT DISTINCT m.chat_id AS id
    FROM message_recipients mr
    JOIN messages m ON mr.message_id = m.id
    WHERE (m.sender_id = 'bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9' OR mr.recipient_id = 'bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9')
) as chat
JOIN  messages m2 ON m2.id = (SELECT id FROM messages AS m3 WHERE m3.chat_id = chat.id ORDER BY m3.created_at DESC LIMIT 1)
JOIN message_recipients mr2 ON mr2.message_id = m2.id
ORDER BY m2.created_at DESC;






DELETE FROM chats;
DELETE FROM messages ;
DELETE FROM message_recipients ;

SELECT * FROM users u;
SELECT DISTINCT chat_id  FROM messages m ;







SELECT * FROM messages m 
WHERE m.chat_id = "72b84e9c-f65e-4835-aa58-2a1f2185fff4"
ORDER BY 
	CASE 
		WHEN m.updated_at is not NULL THEN m.updated_at 
		ELSE m.created_at 
	END
DESC 
	

if request user id = "ini" {
  "sender_id": "ini"
}

{
  "chat_id" : "c70d04a0-7601-4efd-ba32-461d37b65122",
  "sender_id": "bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9",
  "recipient_id": "a8a4960d-1bfe-495a-b53a-1f1325a45188",
  "username":"",
  "last_content"; "hallo",
}

SELECT 
	chat.id AS chat_id, 
	m2.sender_id AS sender_id, 
	m2.content AS last_message, 
	mr2.recipient_id  AS recipient_id, 
	m2.created_at, 
	CASE 
		WHEN mr2.recipient_id = 'bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9' THEN (SELECT username FROM `users` WHERE id =  m2.sender_id) 
		ELSE (SELECT username FROM `users` WHERE id =  mr2.recipient_id) 
	END AS username 
FROM (SELECT DISTINCT m.chat_id AS id FROM message_recipients AS mr 
JOIN messages m ON mr.message_id = m.id 
WHERE (m.sender_id = 'bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9' OR mr.recipient_id = 'bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9')) AS chat 
JOIN messages m2 ON m2.id = (SELECT id FROM messages AS m3 WHERE m3.chat_id = chat.id ORDER BY m3.created_at DESC LIMIT 1) 
JOIN message_recipients mr2 ON mr2.message_id = m2.id 
ORDER BY m2.created_at DESC


SELECT chat.id AS chat_id, m2.sender_id AS sender_id, m2.content AS last_message, 
mr2.recipient_id  AS recipient_id, m2.created_at, 
CASE 
	WHEN mr2.recipient_id = 'bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9' 
		THEN (SELECT username FROM `users` WHERE id =  m2.sender_id) 
	ELSE (SELECT username FROM `users` WHERE id =  mr2.recipient_id) 
	END AS username 
FROM (
	SELECT DISTINCT m.chat_id AS id 
	FROM message_recipients AS mr 
JOIN messages m ON mr.message_id = m.id 
WHERE (m.sender_id = 'bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9' OR mr.recipient_id = 'bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9')) AS chat 
JOIN messages m2 ON m2.id = (SELECT id FROM messages AS m3 WHERE m3.chat_id = chat.id ORDER BY m3.created_at DESC LIMIT 1) 
JOIN message_recipients mr2 ON mr2.message_id = m2.id 
ORDER BY m2.created_at DESC

SELECT chat.id AS chat_id, m2.sender_id AS sender_id, m2.content AS last_message, mr2.recipient_id  AS recipient_id, m2.created_at, 
CASE 
	WHEN mr2.recipient_id = 'bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9' 
		THEN (SELECT username FROM `users` WHERE id =  m2.sender_id) 
	ELSE (SELECT username FROM `users` WHERE id =  mr2.recipient_id) 
	END AS username FROM (SELECT DISTINCT m.chat_id AS id 
FROM message_recipients AS mr 
JOIN messages m ON mr.message_id = m.id WHERE (m.sender_id = 'bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9' OR mr.recipient_id = 'bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9')) AS chat JOIN messages m2 ON m2.id = (SELECT id FROM messages AS m3 WHERE m3.chat_id = chat.id ORDER BY m3.created_at DESC LIMIT 1) 
JOIN message_recipients mr2 ON mr2.message_id = m2.id 
ORDER BY m2.created_at DESC


