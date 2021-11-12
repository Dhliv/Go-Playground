package RandomThings

import "fmt"

type ip struct {
	oct [4]int
}

type rango struct {
	ini, fin ip
}

type red struct {
	ip, mascara, dirRed, broadcast ip
	tamMascara, dispositivos       int
	rango                          rango
}

type redYRango struct {
	ip    ip
	rango rango
	tam   int
}

func (ip ip) impIp() {
	fmt.Printf("%d.%d.%d.%d\n", ip.oct[0], ip.oct[1], ip.oct[2], ip.oct[3])
}

func getMascara(tam int) ip {
	var mascara ip = ip{}

	for i := 0; i < 4 && tam > 0; i++ {
		for j := 7; j >= 0 && tam > 0; j-- {
			mascara.oct[i] |= (1 << j)
			tam--
		}
	}

	return mascara
}

func assignBroadcast(tam int, red *red) {
	for i := 3; i >= 0; i-- {
		if tam > 8 {
			red.broadcast.oct[i] = 255
		} else if tam >= 0 {
			red.broadcast.oct[i] = (1 << (8 - tam)) - 1 + red.dirRed.oct[i]
		} else {
			red.broadcast.oct[i] = red.dirRed.oct[i]
		}
		tam -= 8
	}
}

func sumOne(ip *ip) {
	for k := 3; k >= 0; k-- {
		if ip.oct[k] != 255 {
			ip.oct[k]++
			break
		}
		ip.oct[k] = 0
	}
}

func lessOne(ip *ip) {
	for k := 3; k >= 0; k-- {
		if ip.oct[k] != 0 {
			ip.oct[k]--
			break
		}
		ip.oct[k] = 255
	}
}

func checkRange(ip *ip) {
	var l int

	for i := 3; i >= 0; i-- {
		l = int(ip.oct[i] / 256)

		if l == 0 {
			return
		}
		ip.oct[i] = ip.oct[i] % 256
		ip.oct[i-1] += l
	}
}

func getRanges(tams []int64, rangoAct ip, n int64, rangos *[]redYRango, add bool) ip {
	var ran redYRango

	for i := int64(0); i < n; i++ {
		for j := 2; j <= 32; j++ {
			if tams[i] <= (1<<j)-2 {
				checkRange(&rangoAct)
				ran.ip = rangoAct
				ran.tam = 32 - j

				fmt.Printf("%d.%d.%d.%d/%d and rango:[", rangoAct.oct[0], rangoAct.oct[1], rangoAct.oct[2], rangoAct.oct[3], 32-j)
				var pos int = 3 - int(j/8)
				var rangAuxi = rangoAct
				rangoAct.oct[pos] += 1 << (j % 8)
				var rangAuxf = rangoAct
				lessOne(&rangAuxf)

				ran.rango = rango{ini: rangAuxi, fin: rangAuxf}
				if add {
					*rangos = append(*rangos, ran)
				}

				fmt.Printf("%d.%d.%d.%d -- ", rangAuxi.oct[0], rangAuxi.oct[1], rangAuxi.oct[2], rangAuxi.oct[3])
				fmt.Printf("%d.%d.%d.%d] and Mascara de red: ", rangAuxf.oct[0], rangAuxf.oct[1], rangAuxf.oct[2], rangAuxf.oct[3])
				mask := getMascara(32 - j)
				mask.impIp()
				break
			}
		}
	}

	return rangoAct
}

func redes() {
	var a, b, c, d, e int
	fmt.Scanf("%d.%d.%d.%d/%d", &a, &b, &c, &d, &e)

	var ori ip = ip{oct: [4]int{a, b, c, d}}
	var red red = red{ip: ori,
		mascara:      ip{oct: [4]int{0, 0, 0, 0}},
		tamMascara:   e,
		dispositivos: 32 - e,
	}
	red.mascara = getMascara(e)

	for i := 0; i < 4; i++ {
		red.dirRed.oct[i] = red.ip.oct[i] & red.mascara.oct[i]
	}
	assignBroadcast(red.tamMascara, &red)

	red.rango.ini = red.dirRed
	red.rango.fin = red.broadcast

	var n, size, m int64
	var x, y int
	var rangoAct ip = red.rango.ini
	fmt.Scan(&n, &m)
	var routers, conex []int64
	var ids map[string]int = map[string]int{}
	var name, ax, bx string
	var grafo [][]int
	var names []string
	var rangos []redYRango
	var conexiones [][]int

	for i := int64(0); i < n; i++ {
		fmt.Scan(&size, &name)
		names = append(names, name)
		routers = append(routers, size)
		ids[name] = int(i)
		grafo = append(grafo, make([]int, 0))
	}

	for i := int64(0); i < 2*m; i++ {
		grafo = append(grafo, make([]int, 0))
	}

	for i := int64(0); i < m; i++ {
		fmt.Scan(&size, &ax, &bx)

		if _, ok := ids[ax]; !ok {
			ids[ax] = len(ids)
			names = append(names, ax)
		}

		if _, ok := ids[bx]; !ok {
			ids[bx] = len(ids)
			names = append(names, bx)
		}

		x = ids[ax]
		y = ids[bx]

		conex = append(conex, size)
		conexiones = append(conexiones, make([]int, 2))
		conexiones[i][0] = x
		conexiones[i][1] = y
		grafo[x] = append(grafo[x], y)
		grafo[y] = append(grafo[y], x)
	}

	rangoAct = getRanges(routers, rangoAct, n, &rangos, false)
	rangoAct = getRanges(conex, rangoAct, m, &rangos, true)

	for i := int64(0); i < m; i++ {
		x = conexiones[i][0]
		y = conexiones[i][1]
		fmt.Print("Desde ", names[x], " a ", names[y], " = ")
		var rangeAuxi ip = rangos[i].rango.ini
		var rangeAuxf ip = rangos[i].rango.fin
		sumOne(&rangeAuxi)
		lessOne(&rangeAuxf)
		rangeAuxi.impIp()
		fmt.Print("Desde ", names[y], " a ", names[x], " = ")
		rangeAuxf.impIp()
	}

	fmt.Print("Mascara de red: ")
	red.mascara.impIp()
	fmt.Print("Dirección de red: ")
	red.dirRed.impIp()
	fmt.Print("Broadcast: ")
	red.broadcast.impIp()
}
