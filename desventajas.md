# Desventajas de Go

En este documento voy a ir poniendo las cosas negativas que encuentre acerca de programar en Go.

## Es un infierno de tipos

Después de haber programado por mucho tiempo casi exclusivamente con lenguajes de tipado dinámico (Python, PHP, Elixir), el que Go me requiera ser explicito con todos los tipos que uso me causa conflicto, siento que es demasiado código para expresar pocas instrucciones. A veces ha gastado más tiempo corrigiendo los errores de tipo que en escribir la funcionalidad misma. 

Creo sinceramente que eso se compone con la práctica, te empiezas a acostumbrar a escribir tipos de datos por todos lados y ya no te parece tan grave.

## No inmutable

El que al pasar parámetros a funciones a veces sea por referencia, me parece que hace el código dependiente e inseguro, difícil de razonar en él, aunque quita muchos inconvenientes a la hora de programar.