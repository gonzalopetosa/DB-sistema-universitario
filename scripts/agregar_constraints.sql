alter table alumne 
add constraint alumne_pk 
primary key (id_alumne);

alter table materia 
add constraint materia_pk 
primary key (id_materia);

alter table correlatividad 
add constraint idMateria_fk 
foreign key (id_materia) references materia(id_materia);

alter table correlatividad 
add constraint idMateria_correlativa_fk 
foreign key (id_mat_correlativa) references materia(id_materia);

alter table comision
add constraint comision_pk
primary key (id_materia, id_comision);

alter table cursada
add constraint cursada_pk
primary key (id_materia, id_alumne),
add constraint comision_fk foreign key (id_materia, id_comision) references comision(id_materia, id_comision),
add constraint alumne_fk foreign key (id_alumne) references alumne(id_alumne),
add constraint estado_chk check(estado in ('ingresade', 'aceptade', 'en espera', 'dade de baja'));

alter table periodo
add constraint periodo_pk
primary key (semestre),
add constraint estado_chk check(estado in ('inscripcion', 'cierre inscrip', 'cursada', 'cerrado'));

alter table historia_academica
add constraint historia_academica_pk
primary key (id_alumne, semestre, id_materia),
add constraint id_comision_fk
foreign key (id_materia, id_comision) references comision(id_materia, id_comision),
add constraint estado_chk check(estado in ('ausente', 'reprobada', 'regular', 'aprobada'));

alter table error
add constraint id_error_pk
primary key (id_error),
add constraint operacion_chk check (operacion in ('apertura', 'alta inscrip', 'baja inscrip', 'cierre inscrip', 'aplicacion cupo', 'ingreso nota', 'cierre cursada'));

alter table envio_email
add constraint id_email_pk primary key (id_email),
add constraint estado_chk check (estado in ('pendiente', 'enviado'));

