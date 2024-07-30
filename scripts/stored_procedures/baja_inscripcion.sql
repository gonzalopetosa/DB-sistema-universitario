create or replace function baja_inscripcion(p_id_alumne int, p_id_materia int) returns boolean as $$


declare existe_alumne boolean;
declare semestre_actual char(6);
declare existe_materia boolean;
declare registrado_en_comision int;

begin

    -- Valida que el período exista
    select semestre into semestre_actual from periodo where estado in ('inscripcion', 'cursada');
    if not found then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo) 
        values ('baja inscrip', null, p_id_alumne, p_id_materia, null, current_timestamp, '?no se permiten bajas en este período');
        return false;
    end if;

    -- Validar que el alumne exista
    select count(*) into existe_alumne from alumne where id_alumne = p_id_alumne;
    if not existe_alumne then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo) 
        values ('baja inscrip', semestre_actual, p_id_alumne, p_id_materia, null, current_timestamp, '?id de alumne no válido');
        return false;
    end if;

    -- Validar que la materia exista
   
	select count(*) into existe_materia from materia where id_materia = p_id_materia;
    if not existe_alumne then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo) 
        values ('baja inscrip', semestre_actual, p_id_alumne, p_id_materia, null, current_timestamp, '?id de materia no válido');
        return false;
    end if;

    -- Validar que el alumne esté inscripto en la materia
    select id_comision into registrado_en_comision from cursada where id_alumne = p_id_alumne and id_materia = p_id_materia;
	if not found then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo) 
        values ('baja inscrip', semestre_actual, p_id_alumne, p_id_materia, null, current_timestamp, '?alumne no inscripte en la materia');
        return false;
    end if;

    -- Actualizar la inscripcion
    update cursada set estado = 'dade de baja' where id_alumne = p_id_alumne and id_materia = p_id_materia;

    return true;
end;
$$ language plpgsql;
