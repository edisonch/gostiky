package util

import (
	"github.com/pkg/errors"
	osu "github.com/tredoe/osutil"
	"gostiky/models"
	"strings"
)

func StickIt(st *models.StructSticky) error {
	var err error
	if st.StickyMode > MODE_UNSTICK_APPEND {
		return stickRecursive(st)
	}

	if len(st.Files) < 1 {
		return errors.New("receive no files to process")
	}

	for _, value := range st.Files {
		switch st.StickyMode {
		case MODE_STICK_IMMUTABLE:
			println("MODE_STICK_IMMUTABLE : ", value)
			err = osu.ExecSudo("chattr", "+i",value)
			if err != nil {
				return err
			}
		case MODE_STICK_APPEND:
			println("MODE_STICK_APPEND: ", value)
			err = osu.ExecSudo("chattr", "+a",value)
			if err != nil {
				return err
			}
		case MODE_UNSTICK_IMMUTABLE:
			println("MODE_UNSTICK_IMMUTABLE : ", value)
			err = osu.ExecSudo("chattr", "-i",value)
			if err != nil {
				return err
			}
		case MODE_UNSTICK_APPEND:
			println("MODE_UNSTICK_APPEND : ", value)
			err = osu.ExecSudo("chattr", "-a",value)
			if err != nil {
				return err
			}
		default:
			err = errors.New("Mode Sticky Unknown")
		}
	}
	return err
}

func stickRecursive(st *models.StructSticky) error {
	var err error
	if st.StickyMode < MODE_STICK_IMMUTABLE_RECURSIVE {
		return errors.New("Only Recursive mode can call this")
	}

	if len(st.Directories) < 1 {
		return errors.New("receive no files to process")
	}

	for _, value := range st.Directories {

		if !strings.HasSuffix(value,st.Suffix) {
			continue
		}

		switch st.StickyMode {
		case MODE_STICK_IMMUTABLE_RECURSIVE:
			println("MODE_STICK_IMMUTABLE_RECURSIVE : ", value)
			err = osu.ExecSudo("chattr", "-R","+i"+value)
			if err != nil {
				return err
			}
		case MODE_STICK_APPEND_IMMUTABLE_RECURSIVE:
			println("MODE_STICK_APPEND_IMMUTABLE_RECURSIVE : ", value)
			err = osu.ExecSudo("chattr", "-R","+a"+value)
			if err != nil {
				return err
			}
		case MODE_UNSTICK_IMMUTABLE_RECURSIVE:
			println("MODE_UNSTICK_IMMUTABLE_RECURSIVE : ", value)
			err = osu.ExecSudo("chattr", "-R","-i",value)
			if err != nil {
				return err
			}
		case MODE_UNSTICK_APPEND_IMMUTABLE_RECURSIVE:
			println("MODE_UNSTICK_APPEND_IMMUTABLE_RECURSIVE : ", value)
			err = osu.ExecSudo("chattr", "-R","-a",value)
			if err != nil {
				return err
			}
		default:
			err = errors.New("Mode Sticky Unknown")
		}
	}
	return err
}
