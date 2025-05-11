
 algoritmo "vendas_mensais"
var
    vendas: matriz[12][4] de real
    total_mes: real
    total_ano: real
    i, j: inteiro
inicio
    total_ano := 0

    para i de 0 até 11 faca
        escreva("Digite as vendas do mês ", i+1)
        total_mes := 0
        para j de 0 até 3 faca
            escreva(" Semana ", j+1, ": ")
            leia(vendas[i][j])
            total_mes := total_mes + vendas[i][j]
        fimpara
        escreva("Total do mês ", i+1, ": ", total_mes)
        total_ano := total_ano + total_mes
    fimpara

    escreva("Total de vendas no ano: ", total_ano)
fimalgoritmo
