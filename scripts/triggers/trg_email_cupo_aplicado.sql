drop trigger if exists trg_enviar_email_cupo_aplicado on cursada;

create or replace function enviar_email_cupo_aplicado()
returns trigger as $$
begin
    if new.estado = 'aceptade' and new.nota = null then
        insert into envio_email (f_generacion, email_alumne, asunto, cuerpo, estado)
        values (
            now(),
            (select email from alumne where id_alumne = new.id_alumne),
            'Inscripción aceptada',
            format('Estimade %s %s, su inscripción a la materia %s, comisión %s ha sido aceptada.',
                   (select nombre from alumne where id_alumne = new.id_alumne),
                   (select apellido from alumne where id_alumne = new.id_alumne),
                   (select nombre from materia where id_materia = new.id_materia),
                   new.id_comision),
            'pendiente'
        );
    elsif new.estado = 'en espera' and new.nota = null  then
        insert into envio_email (f_generacion, email_alumne, asunto, cuerpo, estado)
        values (
            now(),
            (select email from alumne where id_alumne = new.id_alumne),
            'Inscripción en espera',
            format('Estimade %s %s, su inscripción a la materia %s, comisión %s está en espera.',
                   (select nombre from alumne where id_alumne = new.id_alumne),
                   (select apellido from alumne where id_alumne = new.id_alumne),
                   (select nombre from materia where id_materia = new.id_materia),
                   new.id_comision),
            'pendiente'
        );
    end if;
    return new;
end;
$$ language plpgsql;

create trigger trg_enviar_email_cupo_aplicado
after update on cursada
for each row
when (new.estado in ('aceptade', 'en espera'))
execute function enviar_email_cupo_aplicado();