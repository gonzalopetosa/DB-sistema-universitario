drop trigger if exists trg_enviar_email_baja_inscripcion on cursada;

create or replace function enviar_email_baja_inscripcion()
returns trigger as $$
begin
	insert into envio_email(f_generacion, email_alumne, asunto, cuerpo, estado)
	values(
		now(),
		(select email from alumne where id_alumne = new.id_alumne),
		'Inscripcion dada de baja',
		format('Estimade %s %s, su inscripcion a la materia %s, comision %s ha sido dada de baja.',
			(select nombre from alumne where id_alumne= new.id_alumne),
			(select apellido from alumne where id_alumne= new.id_alumne),
			(select nombre from materia where id_materia= new.id_materia),
			new.id_comision),
		'pendiente'
	);
	return new;
end;
$$ language plpgsql;

create trigger trg_enviar_email_baja_inscripcion
after update on cursada
for each row
when (new.estado = 'dade de baja')
execute function enviar_email_baja_inscripcion();
