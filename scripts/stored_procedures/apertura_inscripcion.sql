create or replace function apertura_inscripcion(año int, nro_semestre int) returns boolean as $$
declare
    estado_actual char(12);
begin
    -- valida que el año sea mayor o igual al año actual
    if año < extract(year from current_date) then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('apertura', concat(año, '-', nro_semestre), null, null, null, current_timestamp, '?no se permiten inscripciones para un período anterior');
        return false;
    end if;

    -- valida que el número de semestre sea 1 o 2
    if nro_semestre not in (1, 2) then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('apertura', concat(año, '-', nro_semestre), null, null, null, current_timestamp, '?número de semestre no válido');
        return false;
    end if;

    -- valida que el año y semestre solicitado ya exista en la tabla periodo, y que su estado sea cierre inscrip
    select estado into estado_actual from periodo where semestre = concat(año, '-', nro_semestre);
    if found then
        if estado_actual <> 'cierre inscrip' then
            insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
            values ('apertura', concat(año, '-', nro_semestre), null, null, null, current_timestamp, concat('?no es posible reabrir la inscripción del período, estado actual:', estado_actual));
            return false;
        end if;
    end if;

    -- valida que no exista otro período (diferente al solicitado) en estado de inscripcion o cierre inscrip
    if exists (select 1 from periodo where estado in ('inscripcion', 'cierre inscrip') and semestre <> concat(año, '-', nro_semestre)) then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('apertura', concat(año, '-', nro_semestre), null, null, null, current_timestamp, '?no es posible abrir otro período de inscripción');
        return false;
    end if;

    -- inserta o actualiza 
    if found then
        update periodo set estado = 'inscripcion' where semestre = concat(año, '-', nro_semestre);
    else
        insert into periodo (semestre, estado) values (concat(año, '-', nro_semestre), 'inscripcion');
    end if;

    return true;
end;
$$ language plpgsql;
