CREATE OR REPLACE FUNCTION test.user_comment_del(indatauserid integer, indataid integer)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
begin
	delete from test."comments" where id = inDataID and id_user = inDataUserId;
	return;
end;
$function$
;


CREATE OR REPLACE FUNCTION test.user_comment_get(indatauserid integer, indataid integer)
 RETURNS test.comments
 LANGUAGE plpgsql
AS $function$
declare
	outData test."comments"%rowtype;
begin
	select * into outData from test."comments" where id_user = inDataUserId and id = inDataId;
	return outData;
end;
$function$
;

CREATE OR REPLACE FUNCTION test.user_comment_get(indatauserid integer)
    RETURNS SETOF test.comments
    LANGUAGE plpgsql
AS $function$
begin
    return query select * from test."comments" where id_user = inDataUserId;
end;
$function$
;


CREATE OR REPLACE FUNCTION test.user_comment_ins(indatauserid integer, indata character varying)
    RETURNS test.comments
    LANGUAGE plpgsql
AS $function$
declare
    outData test."comments"%rowtype;
begin
    INSERT INTO test."comments"(id_user , txt) select inDataUserId , txt from json_populate_record(null::test."comments",inData::json) returning id, id_user , txt into outData;
    return outData;
end;
$function$
;


CREATE OR REPLACE FUNCTION test.user_comment_upd(indatauserid integer, indataid integer, indata character varying)
    RETURNS test.comments
    LANGUAGE plpgsql
AS $function$
declare
    outData test."comments"%rowtype;
begin
    UPDATE test."comments" SET (txt) = (select txt from json_populate_record(null::test."comments",inData::json))  WHERE id=inDataID and id_user=inDataUserId returning id, id_user , txt  into outData;
    return outData;
end;
$function$
;


CREATE OR REPLACE FUNCTION test.user_del(indataid integer)
    RETURNS void
    LANGUAGE plpgsql
AS $function$
begin
    delete from test.users where id = inDataID;
    return;
end;
$function$
;


CREATE OR REPLACE FUNCTION test.user_get()
    RETURNS SETOF test.users
    LANGUAGE plpgsql
AS $function$
begin
    return query select * from test.users;
end;
$function$
;


CREATE OR REPLACE FUNCTION test.user_get(indata integer)
    RETURNS test.users
    LANGUAGE plpgsql
AS $function$
declare
    outData test.users%rowtype;
begin
    select * into outData from test.users where id = inData;
    return outData;
end;
$function$
;


CREATE OR REPLACE FUNCTION test.user_ins(indata character varying)
    RETURNS test.users
    LANGUAGE plpgsql
AS $function$
declare
    outData test.users%rowtype;
begin
    INSERT INTO test.users("name", email) select "name", email from json_populate_record(null::test.users,inData::json) returning id, "name", email into outData;
    return outData;
end;
$function$
;


CREATE OR REPLACE FUNCTION test.user_upd(indataid integer, indata character varying)
    RETURNS test.users
    LANGUAGE plpgsql
AS $function$
declare
    outData test.users%rowtype;
begin
    UPDATE test.users SET ("name",email) = (select "name", email from json_populate_record(null::test.users,inData::json))  WHERE id=inDataID returning id, "name", email into outData;
    return outData;
end;
$function$
;
