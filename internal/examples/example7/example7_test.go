package example7_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goenum/internal/examples/example7/internal/enums"
)

func TestGameScenario(t *testing.T) {
	require.Equal(t, "게임이 시작되었습니다", processGame("게임-시작"))
	require.Equal(t, "게임이 종료되었습니다", processGame("게임-종료"))
	require.Equal(t, "게임이 일시정지되었습니다", processGame("게임-일시정지"))
	require.Equal(t, "ERROR: INVALID GAME STATE '게임-무효'", processGame("게임-무효"))
}

// processGame demonstrates real-world usage pattern
// processGame 실제 세계의 사용 패턴을 시연
func processGame(input string) string {
	// Step 1: Convert string to enum // 1단계: 문자열을 열거형으로 변환
	gameValue := enums.GameEnum(input)

	// Step 2: Validate enum // 2단계: 열거형 검증
	if !gameValue.Valid() {
		return fmt.Sprintf("ERROR: INVALID GAME STATE '%s'", input)
	}

	// Step 3: Use in switch statement // 3단계: switch문에서 사용
	switch gameValue {
	case enums.Game.S시작():
		return "게임이 시작되었습니다"
	case enums.Game.E종료():
		return "게임이 종료되었습니다"
	case enums.Game.P일시정지():
		return "게임이 일시정지되었습니다"
	default:
		return "알 수 없는 상태"
	}
}
