create or replace function inscripcion_a_materia(p_id_alumne int, p_id_materia int, p_id_comision int) returns boolean as $$
declare
    periodo_actual char(12);
    semestre_actual char(6);
    existe_alumne boolean;
    existe_materia boolean;
    existe_comision boolean;
    esta_inscripto boolean;
    cumple_correlativas boolean;
begin
    -- valida que exista un período en estado de inscripcion
    select estado,semestre into periodo_actual,semestre_actual from periodo where estado = 'inscripcion';
    if not found then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('alta inscrip', null, p_id_alumne, p_id_materia, p_id_comision, current_timestamp, 'período de inscripción cerrado');
        return false;
    end if;

    -- valida que el alumne exista
    select count(*) into existe_alumne from alumne where id_alumne = p_id_alumne;
    if not existe_alumne then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('alta inscrip', semestre_actual, p_id_alumne, p_id_materia, p_id_comision, current_timestamp, 'id de alumne no válido');
        return false;
    end if;

    -- valida que la materia exista
    select count(*) into existe_materia from materia where id_materia = p_id_materia;
    if not existe_materia then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('alta inscrip', semestre_actual, p_id_alumne, p_id_materia, p_id_comision, current_timestamp, 'id de materia no válido');
        return false;
    end if;

    -- valida que la comisión exista para la materia
    select count(*) into existe_comision from comision where id_materia = p_id_materia and id_comision = p_id_comision;
    if not existe_comision then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('alta inscrip', semestre_actual, p_id_alumne, p_id_materia, p_id_comision, current_timestamp, 'id de comisión no válido para la materia');
        return false;
    end if;

    -- valida que le alumne no esté inscripto previamente en la materia (en cualquiera de sus comisiones)
    select count(*) into esta_inscripto from cursada where id_alumne = p_id_alumne and id_materia = p_id_materia;
    if esta_inscripto then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('alta inscrip', semestre_actual, p_id_alumne, p_id_materia, p_id_comision, current_timestamp, 'alumne ya inscripto en la materia');
        return false;
    end if;

    -- valida que le alumne tenga en su historia académica todas las materias correlativas en estado regular o aprobada
    select not exists (
        select 1
        from correlatividad c
        where c.id_materia = p_id_materia
        and not exists (
            select 1
            from historia_academica h
            where h.id_alumne = p_id_alumne
            and h.id_materia = c.id_mat_correlativa
            and h.estado in ('regular', 'aprobada')
        )
    ) into cumple_correlativas;

    if not cumple_correlativas then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('alta inscrip', semestre_actual, p_id_alumne, p_id_materia, p_id_comision, current_timestamp, 'alumne no cumple requisitos de correlatividad');
        return false;
    end if;

    -- inserta la inscripción en la tabla cursada con el estado de ingresade
    insert into cursada (id_materia, id_alumne, id_comision, f_inscripcion, estado)
    values (p_id_materia, p_id_alumne, p_id_comision, current_timestamp, 'ingresade');

    return true;
end;
$$ language plpgsql;
