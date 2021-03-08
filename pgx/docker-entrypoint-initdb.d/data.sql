INSERT INTO iconsfortransactions(mcc, icon)
VALUES (5050, 'icon1'),
       (5000, 'icon2'),
       (5090, 'icon3'),
       (5010, 'icon4')
;

INSERT INTO clients(login, password, full_name, passport, birthday, status)
VALUES ('tolik', 'strongPass-hashed', 'Глухов Анатолий Романович',
        '9214 921041', '1995.02.22', 'ACTIVE'),
       ('mansur', 'bIgPaSs-hashed', 'Иванов Мансур Кириллович',
        '9216 931305', '1999.08.14', 'ACTIVE')
;

INSERT INTO cards(number, balance, issuer, holder, owner_id, status)
VALUES ('-3959', 1491000, 'VISA', 'Глухов Анатолий Романович',
        1, 'ACTIVE'),
       ('-1473', 45370000, 'MIR', 'Иванов Мансур Кириллович',
        2, 'ACTIVE'),
       ('-0312', 9020500, 'VISA', 'Иванов Мансур Кириллович',
        2, 'ACTIVE')
;

INSERT INTO transactions(card_id, sum, mcc, receiver)
VALUES (2, 100000, 5010, '-3959'),
       (2, 100000, 5090, NULL),
       (2, 100000, 5050, 'shop-pay'),
       (2, 500000, 5000, '-1473'),
       (2, 400000, 5050, 'shop-pay'),
       (2, 250000, 5050, 'shop-pay'),
       (2, 1500099, 5010, 'bar')
;

UPDATE cards SET balance = balance + 4700000 WHERE id = 2;



