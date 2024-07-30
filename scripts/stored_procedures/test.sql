create or replace function testear() returns boolean as $$
declare
    rec record;
begin
    for rec in (select * from entrada_trx order by id_orden) loop
        if rec.operacion = 'apertura' then
            perform apertura_inscripcion(rec.año, rec.nro_semestre);
        elsif rec.operacion = 'alta inscrip' then
            perform inscripcion_a_materia(rec.id_alumne, rec.id_materia, rec.id_comision);
        elsif rec.operacion = 'baja inscrip' then
            perform baja_inscripcion(rec.id_alumne, rec.id_materia);
        elsif rec.operacion = 'cierre inscrip' then
            perform cerrar_inscripcion(rec.año, rec.nro_semestre);
        elsif rec.operacion = 'aplicacion cupo' then
            perform aplicar_cupos(rec.año, rec.nro_semestre);
        elsif rec.operacion = 'ingreso nota' then
            perform ingresar_nota_cursada(rec.id_alumne, rec.id_materia, rec.id_comision, rec.nota);
        elsif rec.operacion = 'cierre cursada' then
            perform cerrar_cursada(rec.id_materia, rec.id_comision);
        end if;
    end loop;

    return true;
end;
$$ language plpgsql;
