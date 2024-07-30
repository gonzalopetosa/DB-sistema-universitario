drop trigger if exists trg_enviar_email_inscripcion on cursada;


create or replace function enviar_email_inscripcion()
returns trigger as $$
begin
    insert into envio_email (f_generacion, email_alumne, asunto, cuerpo, estado)
    values (
        now(),
        (select email from alumne where id_alumne = new.id_alumne),
        'Inscripción registrada',
        format('Estimade %s %s, su inscripción a la materia %s, comisión %s ha sido registrada exitosamente.',
               (select nombre from alumne where id_alumne = new.id_alumne),
               (select apellido from alumne where id_alumne = new.id_alumne),
               (select nombre from materia where id_materia = new.id_materia),
               new.id_comision),
        'pendiente'
    );
    return new;
end;
$$ language plpgsql;


create trigger trg_enviar_email_inscripcion
after insert on cursada
for each row
when (new.estado = 'ingresade')
execute function enviar_email_inscripcion();