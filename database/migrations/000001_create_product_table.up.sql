CREATE TABLE feira (
    id         SERIAL PRIMARY KEY NOT NULL,
  	long       VARCHAR(50) NOT NULL,  	
	lat        VARCHAR(50) NOT NULL,  	
	setcens    VARCHAR(50) NOT NULL,	
	areap      VARCHAR(50) NOT NULL,	
	coddist    VARCHAR(50) NOT NULL,	
	distrito   VARCHAR(50) NOT NULL,	
	codsubpref VARCHAR(50) NOT NULL,	
	subprere   VARCHAR(50) NOT NULL,	
	regiao5    VARCHAR(50) NOT NULL,	
	regiao8    VARCHAR(50) NOT NULL,	
	nomefreira VARCHAR(50) NOT NULL,	
	registo    VARCHAR(50) NOT NULL,	
	logradouro VARCHAR(50) NOT NULL,	
	numero     VARCHAR(50) NOT NULL,	
	bairro     VARCHAR(50) NOT NULL,	
	referencia VARCHAR(50) NOT NULL	  
);
CREATE UNIQUE INDEX nomefreira_idx ON public.feira (nomefreira);