create or replace function ingresar_nota_cursada(p_id_alumne int, p_id_materia int, p_id_comision int, p_nota int)  returns boolean as $$
declare
    estado_periodo char(12);
    semestre_actual char(6);
    existe_alumno boolean;
    existe_materia boolean;
    existe_comision boolean;
    existe_inscripcion boolean;
begin
    -- valida que el periodo esté en 'cursada'
    select estado, semestre into estado_periodo, semestre_actual
    from periodo
    where estado = 'cursada';
    if estado_periodo is null then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('ingreso nota', semestre_actual, p_id_alumne, p_id_materia, p_id_comision, now(), '?período de cursada cerrado');
        return false;
    end if;

    -- valida que el alumne exista
    select count(*) > 0 into existe_alumno
    from alumne
    where id_alumne = p_id_alumne;

    if not existe_alumno then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('ingreso nota', semestre_actual, p_id_alumne, p_id_materia, p_id_comision, now(), '?id de alumne no válido');
        return false;
    end if;

    -- valida que la materia exista
    select count(*) > 0 into existe_materia
    from materia
    where id_materia = p_id_materia;

    if not existe_materia then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('ingreso nota', semestre_actual, p_id_alumne, p_id_materia, p_id_comision, now(), '?id de materia no válido');
        return false;
    end if;

    -- valida que la comision exista para la materia
    select count(*) > 0 into existe_comision
    from comision
    where id_materia = p_id_materia and id_comision = p_id_comision;

    if not existe_comision then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('ingreso nota', semestre_actual, p_id_alumne, p_id_materia, p_id_comision, now(), '?id de comisión no válido para la materia');
        return false;
    end if;

    -- valida que el alumne esté inscripte
    select count(*) > 0 into existe_inscripcion
    from cursada
    where id_alumne = p_id_alumne and id_materia = p_id_materia and id_comision = p_id_comision and estado = 'aceptade';

    if not existe_inscripcion then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('ingreso nota', semestre_actual, p_id_alumne, p_id_materia, p_id_comision, now(), '?alumne no cursa en la comisión');
        return false;
    end if;

    -- validar que la nota esté en el rango de 0 a 10
    if p_nota < 0 or p_nota > 10 then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('ingreso nota', semestre_actual, p_id_alumne, p_id_materia, p_id_comision, now(), '?nota no válida:' || p_nota);
        return false;
    end if;

    -- actualiza la nota del alumne
    update cursada
    set nota = p_nota
    where id_alumne = p_id_alumne and id_materia = p_id_materia and id_comision = p_id_comision;

    return true;
end;
$$ language plpgsql;
