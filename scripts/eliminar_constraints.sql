alter table cursada
drop constraint cursada_pk,
drop constraint comision_fk, 
drop constraint alumne_fk, 
drop constraint estado_chk;

alter table periodo
drop constraint periodo_pk,
drop constraint estado_chk;

alter table correlatividad 
drop constraint idMateria_fk, 
drop constraint idMateria_correlativa_fk;

alter table historia_academica
drop constraint historia_academica_pk,
drop constraint id_comision_fk,
drop constraint estado_chk;

alter table error
drop constraint id_error_pk,
drop constraint operacion_chk;

alter table envio_email
drop constraint id_email_pk,
drop constraint estado_chk;

alter table alumne 
drop constraint alumne_pk;

alter table materia 
drop constraint materia_pk;

alter table comision
drop constraint comision_pk;