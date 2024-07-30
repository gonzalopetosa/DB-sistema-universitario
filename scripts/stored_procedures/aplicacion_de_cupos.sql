create or replace function aplicar_cupos(p_año int, p_semestre int) returns boolean as $$
declare
    periodo_actual char(15);
    cupo_comision int;
    registro record;
begin
    -- semestre en estado cierre inscrip
    select estado into periodo_actual from periodo where semestre = concat(p_año, '-', p_semestre) and estado = 'cierre inscrip';
    if not found then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('aplicacion cupo', concat(p_año, '-', p_semestre), null, null, null, current_timestamp, '?el semestre no se encuentra en un período válido para aplicar cupos');
        return false;
    end if;

    for registro in (select id_materia, id_comision, cupo from comision) loop
        cupo_comision := registro.cupo;

        -- cambia estado a aceptade
        update cursada
        set estado = 'aceptade'
        where (id_alumne, id_materia) in (
            select id_alumne, id_materia
            from (
                select id_alumne, id_materia
                from cursada
                where id_materia = registro.id_materia
                  and id_comision = registro.id_comision
                  and estado = 'ingresade'
                order by f_inscripcion
                limit registro.cupo
            ) as subquery
        );

        -- los ingresade pasan a en espera
        update cursada
        set estado = 'en espera'
        where id_materia = registro.id_materia
          and id_comision = registro.id_comision
          and estado = 'ingresade';
    end loop;

    -- acá pasa a 'cursada' el periodo
    update periodo
    set estado = 'cursada'
    where semestre = concat(p_año, '-', p_semestre);

    return true;
exception
    when others then
        insert into error (operacion, semestre, id_alumne, id_materia, id_comision, f_error, motivo)
        values ('aplicacion cupo', concat(p_año, '-', p_semestre), null, null, null, current_timestamp, 'error durante la aplicación de cupos');
        return false;
end;
$$ language plpgsql;
