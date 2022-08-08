package gl

import (
	"guillermoSb/glLibrary/numg"
)
type KWA map[string]interface{}

func flatShader(r *Renderer, args KWA) (float32, float32, float32)  {
	// u := args["baryCoords"].(V3).X
	// v := args["baryCoords"].(V3).Y
	// w := args["baryCoords"].(V3).Z
	triangleNormal := args["triangleNormal"].(V3)
	color := args["vColor"].(Color)
	A := numg.V3{triangleNormal.X, triangleNormal.Y, triangleNormal.Z}
	B := numg.V3MultiplyScalar(numg.V3{r.dirLight.X, r.dirLight.Y, r.dirLight.Z}, -1)
	
	intensity := numg.V3DotProduct(A,B)

	if intensity > 0 {
			return color.R * intensity,color.B * intensity, color.G * intensity		
	} else {
			return 0,0,0
	}
	
}

func textureShader(r *Renderer, args KWA) (float32, float32, float32)  {
	u := args["baryCoords"].(V3).X
	v := args["baryCoords"].(V3).Y
	w := args["baryCoords"].(V3).Z
	// triangleNormal := args["triangleNormal"].(V3)
	color := args["vColor"].(Color)
	textCoords := args["textCoords"].([][]float32)
	tA := textCoords[0]
	tB := textCoords[1]
	tC := textCoords[2]
	if r.UseTexture {
		// p = Au + Bv + Cw
		tU := tA[0] * u + tB[0] * v + tC[0] * w
		tV := tA[1] * u + tB[1] * v + tC[1] * w
		texColor, err := r.activeTexture.GetColor(tU, tV)
		if err == nil {
			color.B *= texColor.B
			color.R *= texColor.R
			color.G *= texColor.G
		}
	}
	return color.R,color.G,color.B
}