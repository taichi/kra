CREATE TABLE films (
    code        char(5) PRIMARY KEY,
    title       varchar(40) NOT NULL,
    did         integer NOT NULL,
    date_prod   date,
    kind        varchar(10),
    len         interval hour to minute
);

INSERT INTO films (code, title, did, date_prod, kind, len) VALUES ('1111', 'aaaa', 12, CURRENT_DATE, 'CDR', INTERVAL '01:30:00');
INSERT INTO films (code, title, did, date_prod, kind, len) VALUES ('2222', 'bbbb', 34, CURRENT_DATE, 'ZDE', INTERVAL '02:00:00');
INSERT INTO films (code, title, did, date_prod, kind, len) VALUES ('3333', 'cccc', 52, CURRENT_DATE, 'IOM', INTERVAL '01:00:00');
INSERT INTO films (code, title, did, date_prod, kind, len) VALUES ('4444', 'dddd', 76, CURRENT_DATE, 'ERW', INTERVAL '01:15:00');
