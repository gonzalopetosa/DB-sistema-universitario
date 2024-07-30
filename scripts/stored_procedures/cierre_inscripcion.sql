create or replace function cerrar_inscripcion(p_año int, p_nro_semestre int) returns boolean as $$
declare
    periodo_actual char(15);
begin

    -- Valida que el periodo esté en 'inscripcion'
    select estado into periodo_actual from periodo where semestre = concat(p_año, '-', p_nro_semestre) and estado = 'inscripcion';
    if not found then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('cierre inscrip', concat(p_año, '-', p_nro_semestre), null, null, null, current_timestamp, '?el semestre no se encuentra en período de inscripción');
        return false;
    end if;

    -- Cambia el periodo a 'cierre inscrip'
    update periodo
    set estado = 'cierre inscrip'
    where semestre = concat(p_año, '-', p_nro_semestre);

    return true;
end;
$$ language plpgsql;
