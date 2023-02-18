package uuid

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBase64(t *testing.T) {
	u, err := NewV4()
	require.NoError(t, err)
	require.NotEmpty(t, u)

	b64 := u.Base64Url()
	require.Len(t, b64, 22)

	u2, err := FromBase64Url(b64)
	require.NoError(t, err)
	require.Equal(t, u, u2)
}

func TestFromBase64(t *testing.T) {
	u, err := FromBase64Url("mtmLveavS76ZAqWI4lC8sA")
	require.NoError(t, err)
	require.NotEmpty(t, u)
	require.Equal(t, "9ad98bbd-e6af-4bbe-9902-a588e250bcb0", u.String())
}

func TestV7(t *testing.T) {
	u, _ := NewV7()
	time.Sleep(10 * time.Millisecond)
	u2, _ := NewV7()
	require.Less(t, u.Base64Url(), u2.Base64Url())
}

func TestProquint(t *testing.T) {
	u, err := NewV4()
	require.NoError(t, err)
	require.NotEmpty(t, u)

	p := u.Proquint()
	require.Len(t, p, 47)

	u2, err := FromProquint(p)
	require.NoError(t, err)
	require.Equal(t, u, u2)
}

func TestFromProquint(t *testing.T) {
	u, err := FromProquint("ginut-sasof-rujus-hodug-nisaz-dafig-fajan-puvoh")
	require.NoError(t, err)
	require.Equal(t, "367dc322-bd7c-4873-970f-10932149afa4", u.String())
}
