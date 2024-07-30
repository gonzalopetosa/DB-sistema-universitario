drop trigger if exists trg_actualizar_estado_en_espera on cursada;

create or replace function actualizar_estado_en_espera() returns trigger as $$
declare
    alumne_en_espera int;
begin
    -- si el nuevo estado es 'dade de baja'
    if new.estado = 'dade de baja' then
        
        if exists (select 1 from periodo where estado = 'cursada') then
            -- selecciono el primer alumne en espera
            select id_alumne into alumne_en_espera
            from cursada
            where id_materia = new.id_materia
              and id_comision = new.id_comision
              and estado = 'en espera'
            order by f_inscripcion asc
            limit 1;

            -- actualizo el estado
            if alumne_en_espera is not null then
                update cursada
                set estado = 'aceptade'
                where id_alumne = alumne_en_espera;
            end if;
        end if;
    end if;

    return new;
end;
$$ language plpgsql;

create trigger trg_actualizar_estado_en_espera after update on cursada
for each row
when (new.estado = 'dade de baja')
execute function actualizar_estado_en_espera();
