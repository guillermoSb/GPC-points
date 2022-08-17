package gl

import (
	"guillermoSb/glLibrary/numg"
	"math"
)
type KWA map[string]interface{}


func BoxBlurShader(r *Renderer, args KWA) (float32, float32, float32) {
	u := args["baryCoords"].(V3).X
	v := args["baryCoords"].(V3).Y
	w := args["baryCoords"].(V3).Z
	triangleNormal := args["triangleNormal"].(V3)
	color := args["vColor"].(Color)
	A := numg.V3{triangleNormal.X, triangleNormal.Y, triangleNormal.Z}
	B := numg.V3MultiplyScalar(numg.V3{r.dirLight.X, r.dirLight.Y, r.dirLight.Z}, -1)
	intensity := numg.V3DotProduct(A,B)
	textCoords := args["textCoords"].([][]float32)
	tA := textCoords[0]
	tB := textCoords[1]
	tC := textCoords[2]
	if r.ActiveTexture.Name != "" {
		// p = Au + Bv + Cw
		tU := tA[0] * u + tB[0] * v + tC[0] * w
		tV := tA[1] * u + tB[1] * v + tC[1] * w
		texColor, err := r.ActiveTexture.GetColor(tU, tV)
		texColor2, err := r.ActiveTexture.GetColor(tU + 0.01, tV + 0.01)
		texColor3, err := r.ActiveTexture.GetColor(tU - 0.01, tV - 0.01)
		texColor4, err := r.ActiveTexture.GetColor(tU, tV - 0.01)
		texColor5, err := r.ActiveTexture.GetColor(tU, tV + 0.01)
		if err == nil {
			color.B *= (texColor.B + texColor2.B + texColor3.B + texColor4.B + texColor5.B) / 5
			color.R *= (texColor.R + texColor2.R + texColor3.R + texColor4.R + texColor5.R) / 5
			color.G *= (texColor.G + texColor2.G + texColor3.G + texColor4.G + texColor5.G) / 5
		}
		
		if intensity > 0 {
			return color.R * intensity,color.G * intensity, color.B * intensity
		} else {
			return 0,0,0
		}
	}
		return color.R ,color.G, color.B 
}

func FlatShader(r *Renderer, args KWA) (float32, float32, float32)  {
	u := args["baryCoords"].(V3).X
	v := args["baryCoords"].(V3).Y
	w := args["baryCoords"].(V3).Z
	triangleNormal := args["triangleNormal"].(V3)
	color := args["vColor"].(Color)
	A := numg.V3{triangleNormal.X, triangleNormal.Y, triangleNormal.Z}
	B := numg.V3MultiplyScalar(numg.V3{r.dirLight.X, r.dirLight.Y, r.dirLight.Z}, -1)
	intensity := numg.V3DotProduct(A,B)
	textCoords := args["textCoords"].([][]float32)
	tA := textCoords[0]
	tB := textCoords[1]
	tC := textCoords[2]
	if r.ActiveTexture.Name != "" {
		// p = Au + Bv + Cw
		tU := tA[0] * u + tB[0] * v + tC[0] * w
		tV := tA[1] * u + tB[1] * v + tC[1] * w
		texColor, err := r.ActiveTexture.GetColor(tU, tV)
		if err == nil {
			color.B *= texColor.B
			color.R *= texColor.R
			color.G *= texColor.G
		}
			if intensity > 0 {
			return color.R * intensity,color.G * intensity, color.B * intensity		
			} else {
			return 0,0,0
		}
	}
		return color.R ,color.G, color.B 
}

func UnlitShader(r *Renderer, args KWA) (float32, float32, float32)  {
	u := args["baryCoords"].(V3).X
	v := args["baryCoords"].(V3).Y
	w := args["baryCoords"].(V3).Z
	// triangleNormal := args["triangleNormal"].(V3)
	color := args["vColor"].(Color)
	textCoords := args["textCoords"].([][]float32)
	tA := textCoords[0]
	tB := textCoords[1]
	tC := textCoords[2]
	if r.ActiveTexture.Name != "" {
		// p = Au + Bv + Cw
		tU := tA[0] * u + tB[0] * v + tC[0] * w
		tV := tA[1] * u + tB[1] * v + tC[1] * w
		texColor, err := r.ActiveTexture.GetColor(tU, tV)
		if err == nil {
			color.B *= texColor.B
			color.R *= texColor.R
			color.G *= texColor.G
		}
	}
	return color.R,color.G,color.B
}

func GShader(r *Renderer, args KWA) (float32, float32, float32)  {
	u := args["baryCoords"].(V3).X
	v := args["baryCoords"].(V3).Y
	w := args["baryCoords"].(V3).Z
	x := args["positionCoords"].(V2).X
	// y := args["positionCoords"].(V2).Y
	
	// triangleNormal := args["triangleNormal"].(V3)
	color := args["vColor"].(Color)
	textCoords := args["textCoords"].([][]float32)
	tA := textCoords[0]
	tB := textCoords[1]
	tC := textCoords[2]
	if r.ActiveTexture.Name != "" {
		// p = Au + Bv + Cw
		tU := tA[0] * u + tB[0] * v + tC[0] * w
		tV := tA[1] * u + tB[1] * v + tC[1] * w
		texColor, err := r.ActiveTexture.GetColor(tU, tV)
		if err == nil {
			color.B *= float32(math.Max(float64(0.5*(float32(math.Cos(float64(x)))+1) * texColor.B), 0))
			color.R *= float32(math.Max(float64(0.5*(float32(math.Cos(float64(x)))+1) * texColor.R), 0))
			color.G *= float32(math.Max(float64(0.5*(float32(math.Cos(float64(x)))+1) * texColor.G), 0))
		}
	}
	return color.R,color.G,color.B
}