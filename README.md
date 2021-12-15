# Tarea 3 Sistemas Distribuidos - Replicación

- Álvaro Fuentes (Par. 200) 201611539-0
- Christian Trujillo (Par. 201) 201673582-8

Versión GO: go1.17.3 linux/amd64 (ya esta actualizado en VM's)

Para ejecutar los procesos, en cada maquina: 
- Navegar a carpeta /SD-T3
- Ejecutar comando make dependiendo del tipo de proceso

Formato: Proceso: Maquina - IP 

## Leia: dist113 - 10.6.43.101

```
make leia
```

## Fulcrum x: dist113 - 10.6.43.101

```
make fulcrum
```

## Ahsoka: dist114 -10.6.43.102

```
make info1
```

##  Fulcrum y: dist114 - 10.6.43.102

```
make fulcrum
```

## Thrawn:  dist115 -10.6.43.103

```
make info2
```

## Fulcrum z:  dist115 -10.6.43.103

```
make fulcrum
```

## broker Mos Eisley: dist116 -10.6.43.104

```
make broker
```

## Ejemplo de consulta que cumple Read Your Writes

- Informante crea planeta1 ciudad1 valor1
- Broker le dice que lo haga en fulcrum1 (aleatoriamente elegido)
- Informante lo hace y broker guarda fulcrum1 como last_fulcrum para ese planeta
- Fulcrum1 escribe (1,0,0) en el vector planeta1 y deja el comando en planeta1-log.txt
- Posteriormente informante quiere editar ciudad1 y le avisa al broker (enviado planeta, ciudad, ultimo IP visitada y vector del planeta)
- Si la última ip de Fulcrum visitada no es vacía, se conecta a esa ip
- Busca el vector del planeta correspondiente
- Si no existe dicho vector o  existe y el vector del planeta otorgado por el informate es menor al vector de Fulcrum (en la posición correspondiente), se busca otro servidor Fulcrum con el vector más reciente en dicha posición
- En caso contrario se opera sobre la última IP
- Informante se conecta a fulcrum1 (por ejemplo) y edita
- Fulcrum1 escribe (2,0,0) en el vector planeta1 y continúa...

## Ejemplo de consulta que cumple Monotonic Reads

- Leia consulta por planeta1 ciudad1
- Broker se fija los vectores de planeta1
- Broker elige el vector de planeta1 con mas cambios (en teoría el mas actualizado)
- Broker envia la consulta de leia al fulcrum1 (x ej)
- Fulcrum1 responde al broker
- Broker le entrega la respuesta a Leia ademas del vector del planeta consultado
- Leia guarda en memoria la consulta que hizo, junto con su vector y el fulcrum que le respondio por ultima vez cuando pregunto por ese planeta
- La proxima vez Leia enviará estos últimos datos al broker y los comparara con los actuales del fulcrum, si no hay cambios responderá con lo que tenia ya guardado
