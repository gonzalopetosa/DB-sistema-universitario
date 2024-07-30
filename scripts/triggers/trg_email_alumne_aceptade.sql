drop trigger if exists trg_enviar_email_aceptacion_lista_espera on cursada;

-- Función para enviar email cuando un alumno sale de la lista de espera y es aceptado
create or replace function enviar_email_aceptacion_lista_espera()
returns trigger as $$
begin
    insert into envio_email (f_generacion, email_alumne, asunto, cuerpo, estado)
    values (
        now(),
        (select email from alumne where id_alumne = new.id_alumne),
        'Inscripción aceptada',
        format('Estimade %s %s, su inscripción a la materia %s, comisión %s ha sido aceptada desde la lista de espera.',
               (select nombre from alumne where id_alumne = new.id_alumne),
               (select apellido from alumne where id_alumne = new.id_alumne),
               (select nombre from materia where id_materia = new.id_materia),
               new.id_comision),
        'pendiente'
    );
    return new;
end;
$$ language plpgsql;

create trigger trg_enviar_email_aceptacion_lista_espera after update on cursada
for each row
when (old.estado = 'en espera' and new.estado = 'aceptade')
execute function enviar_email_aceptacion_lista_espera();
