create table if not exists alumne(
	id_alumne int, 
	nombre text, 
	apellido text, 
	dni int, 
	fecha_nacimiento date, 
	telefono char(12), 
	email text
);

create table if not exists materia(
	id_materia int, 
	nombre text
);

create table if not exists correlatividad(
	id_materia int,
	id_mat_correlativa int
);


create table if not exists historia_academica(
    id_alumne int,
    semestre text,
    id_materia int,
    id_comision int,
    estado char(15),
    nota_regular int,
    nota_final int
);

create table if not exists error(
    id_error serial,
    operacion char(15),
    semestre text,
    id_alumne int,
    id_materia int,
    id_comision int,
    f_error timestamp,
    motivo varchar(80)
);

create table if not exists envio_email(
    id_email serial,
    f_generacion timestamp,
    email_alumne text,
    asunto text,
    cuerpo text,
    f_envio text,
    estado char(10)    
);

create table if not exists entrada_trx (
    id_orden int,
    operacion char(15),
    año int,
    nro_semestre int,
    id_alumne int,
    id_materia int,
    id_comision int,
    nota int
);


create table if not exists comision(
	id_materia int,
	id_comision int,
	cupo int
);

create table if not exists cursada(
	id_materia int,
	id_alumne int,
	id_comision int,
	f_inscripcion timestamp,
	nota int,
	estado char(12)
);

create table if not exists periodo(
	semestre text,
	estado char(15)
);


