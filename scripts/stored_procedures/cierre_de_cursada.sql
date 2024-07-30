create or replace function cerrar_cursada(p_id_materia int, p_id_comision int ) returns boolean as $$
declare
    v_estado_periodo char(12);
    semestre_actual char(6);
    existe_materia boolean;
    existe_comision boolean;
    v_alumnes_inscriptes int;
    notas_completas boolean;
begin

    select estado, semestre into v_estado_periodo, semestre_actual
    from periodo
    where estado = 'cursada';

    if v_estado_periodo is null then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('cierre cursada', semestre_actual, null, p_id_materia, p_id_comision, now(), 'período de cursada cerrado');
        return false;
    end if;

    select count(*) > 0 into existe_materia
    from materia
    where id_materia = p_id_materia;

    if not existe_materia then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('cierre cursada', semestre_actual, null, p_id_materia, p_id_comision, now(), 'id de materia no válido');
        return false;
    end if;

    select count(*) > 0 into existe_comision
    from comision
    where id_materia = p_id_materia and id_comision = p_id_comision;

    if not existe_comision then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('cierre cursada', semestre_actual, null, p_id_materia, p_id_comision, now(), 'id de comisión no válido para la materia');
        return false;
    end if;

    select count(*) into v_alumnes_inscriptes
    from cursada
    where id_materia = p_id_materia and id_comision = p_id_comision;

    if v_alumnes_inscriptes = 0 then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('cierre cursada', semestre_actual, null, p_id_materia, p_id_comision, now(), 'comisión sin alumnes inscriptes');
        return false;
    end if;

    select count(*) = 0 into notas_completas
    from cursada
    where id_materia = p_id_materia and id_comision = p_id_comision and estado = 'aceptade' and nota is null;

    if not notas_completas then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('cierre cursada', semestre_actual, null, p_id_materia, p_id_comision, now(), 'la carga de notas no está completa');
        return false;
    end if;

    insert into historia_academica (id_alumne, semestre, id_materia, id_comision, estado, nota_regular, nota_final)
    select id_alumne, semestre_actual, id_materia, id_comision,
        case
            when nota = 0 then 'ausente'
            when nota between 1 and 3 then 'reprobada'
            when nota between 4 and 6 then 'regular'
            when nota between 7 and 10 then 'aprobada'
        end as estado,
        nota as nota_regular,
        case
            when nota between 7 and 10 then nota
            else null
        end as nota_final
    from cursada
    where id_materia = p_id_materia and id_comision = p_id_comision and estado = 'aceptade';

    -- raise exception 'error personalizado antes de eliminar registros de cursada';

    delete from cursada
    where id_materia = p_id_materia and id_comision = p_id_comision;

    update periodo set estado = 'cerrado' where estado = 'cursada';

    return true;
end;
$$ language plpgsql;
