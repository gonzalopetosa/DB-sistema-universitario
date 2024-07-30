drop trigger if exists trg_cierre_cursada on historia_academica;


create or replace function fn_cierre_cursada() returns trigger as $$
declare
    semestre_actual char(6);
    v_nombre_alumne text;
    v_apellido_alumne text;
    v_email_alumne text;
    v_nombre_materia text;
    v_id_comision int;
    v_estado_academico char(15);
    v_nota_regular int;
    v_nota_final int;
begin

    select semestre into semestre_actual from periodo 
    order by semestre desc
    limit 1;

    -- datos del alumno
    select nombre, apellido, email into v_nombre_alumne, v_apellido_alumne, v_email_alumne
    from alumne
    where id_alumne = new.id_alumne;

    -- datos de la materia y comisión
    select nombre into v_nombre_materia
    from materia
    where id_materia = new.id_materia;
    v_id_comision = new.id_comision;

    -- estado académico y notas
    select estado, nota_regular, nota_final into v_estado_academico, v_nota_regular, v_nota_final
    from historia_academica
    where id_alumne = new.id_alumne and id_materia = new.id_materia and id_comision = new.id_comision and semestre = semestre_actual;

    -- inserto el email
    insert into envio_email (f_generacion, email_alumne, asunto, cuerpo, estado)
    values (
        now(),
        v_email_alumne,
        'cierre de cursada',
        concat(
            'estimado/a ', v_nombre_alumne, ' ', v_apellido_alumne, ', la cursada de la materia ',
            v_nombre_materia, ' en la comisión ', v_id_comision, ' ha sido cerrada. Su estado académico es ',
            v_estado_academico, ' y su nota es ',
            case when v_estado_academico = 'aprobada' then v_nota_final else v_nota_regular end, '.'
        ),
        'pendiente'
    );

    return new;
end;
$$ language plpgsql;

-- agrego definición del trigger!!
create trigger trg_cierre_cursada
after insert on historia_academica
for each row
execute function fn_cierre_cursada();
