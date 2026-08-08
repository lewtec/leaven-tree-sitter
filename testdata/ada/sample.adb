package P is
   type A is access Integer;
   type B is access not null Integer;
   type C is access constant Integer;
   type D is access all Integer;
   type E is access function return Boolean;
   type F is access protected function return Boolean;
end;
