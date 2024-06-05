package cli

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Flag struct {
	// Long name of the flag.
	Long string
	// Short name of the flag (one character).
	Short string
	// Usage is the description shown in the 'help' output.
	Usage string
	// Required is used to mark the flag as required.
	Required bool
	// Persistent is used to propagate the flag to subcommands.
	Persistent bool
	// Default is the default value when the flag is not explicitly supplied. It should have the same type as the value
	// behind the pointer in field Ptr.
	Default interface{}
	// Ptr is a pointer to the value into which the flag will be parsed.
	Ptr interface{}
	// Hidden is used to mark the flag as hidden.
	Hidden bool
}

func BuildFlags(obj interface{}) []Flag {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr {
		panic(fmt.Errorf("expected a pointer, got %s", v.Kind()))
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("expected a struct, got %s", v.Kind()))
	}
	t := v.Type()

	var err error
	flags := make([]Flag, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		flags[i], err = buildFlag(v.Field(i), t.Field(i))
		if err != nil {
			panic(err)
		}
	}
	return flags
}

func buildFlag(val reflect.Value, sf reflect.StructField) (Flag, error) {
	const (
		tagNameLong       = "long"
		tagNameShort      = "short"
		tagNameRequired   = "required"
		tagNamePersistent = "persistent"
		tagNameUsage      = "usage"
		tagNameHidden     = "hidden"
	)

	var (
		long       string
		short      string
		required   bool
		persistent bool
		usage      string
		hidden     bool
	)

	if v, ok := sf.Tag.Lookup(tagNameLong); ok {
		long = v
	}
	if v, ok := sf.Tag.Lookup(tagNameShort); ok {
		short = v
	}
	if v, ok := sf.Tag.Lookup(tagNameRequired); ok {
		var err error
		required, err = strconv.ParseBool(v)
		if err != nil {
			return Flag{}, fmt.Errorf("error parsing tag \"required\": %w", err)
		}
	}
	if v, ok := sf.Tag.Lookup(tagNamePersistent); ok {
		var err error
		persistent, err = strconv.ParseBool(v)
		if err != nil {
			return Flag{}, fmt.Errorf("error parsing tag \"persistent\": %w", err)
		}
	}
	if v, ok := sf.Tag.Lookup(tagNameUsage); ok {
		usage = v
	}
	if v, ok := sf.Tag.Lookup(tagNameHidden); ok {
		var err error
		hidden, err = strconv.ParseBool(v)
		if err != nil {
			return Flag{}, fmt.Errorf("error parsing tag \"hidden\": %w", err)
		}
	}

	return Flag{
		Long:       long,
		Short:      short,
		Usage:      usage,
		Required:   required,
		Persistent: persistent,
		Default:    nil,
		Ptr:        val.Addr().Interface(),
		Hidden:     hidden,
	}, nil
}

func setFlagsToCommand(cmd *cobra.Command, flags []Flag) {
	for _, f := range flags {
		var flags *pflag.FlagSet
		if f.Persistent {
			flags = cmd.PersistentFlags()
		} else {
			flags = cmd.Flags()
		}

		if f.Required {
			f.Usage += " (required)"
		}

		switch val := f.Ptr.(type) {
		case *string:
			if f.Default == nil {
				f.Default = ""
			}
			flags.StringVarP(val, f.Long, f.Short, f.Default.(string), f.Usage)
		case *int:
			if f.Default == nil {
				f.Default = 0
			}
			flags.IntVarP(val, f.Long, f.Short, f.Default.(int), f.Usage)
		case *int8:
			if f.Default == nil {
				f.Default = int8(0)
			}
			flags.Int8VarP(val, f.Long, f.Short, f.Default.(int8), f.Usage)
		case *int16:
			if f.Default == nil {
				f.Default = int16(0)
			}
			flags.Int16VarP(val, f.Long, f.Short, f.Default.(int16), f.Usage)
		case *int32:
			if f.Default == nil {
				f.Default = int32(0)
			}
			flags.Int32VarP(val, f.Long, f.Short, f.Default.(int32), f.Usage)
		case *int64:
			if f.Default == nil {
				f.Default = int64(0)
			}
			flags.Int64VarP(val, f.Long, f.Short, f.Default.(int64), f.Usage)
		case *float32:
			if f.Default == nil {
				f.Default = float32(0)
			}
			flags.Float32VarP(val, f.Long, f.Short, f.Default.(float32), f.Usage)
		case *float64:
			if f.Default == nil {
				f.Default = float64(0)
			}
			flags.Float64VarP(val, f.Long, f.Short, f.Default.(float64), f.Usage)
		case *bool:
			if f.Default == nil {
				f.Default = false
			}
			flags.BoolVarP(val, f.Long, f.Short, f.Default.(bool), f.Usage)
		case *time.Duration:
			if f.Default == nil {
				f.Default = time.Duration(0)
			}
			flags.DurationVarP(val, f.Long, f.Short, f.Default.(time.Duration), f.Usage)
		case *[]bool:
			if f.Default == nil {
				f.Default = []bool(nil)
			}
			flags.BoolSliceVarP(val, f.Long, f.Short, f.Default.([]bool), f.Usage)
		case *[]float32:
			if f.Default == nil {
				f.Default = []float32(nil)
			}
			flags.Float32SliceVarP(val, f.Long, f.Short, f.Default.([]float32), f.Usage)
		case *[]float64:
			if f.Default == nil {
				f.Default = []float64(nil)
			}
			flags.Float64SliceVarP(val, f.Long, f.Short, f.Default.([]float64), f.Usage)
		case *[]int32:
			if f.Default == nil {
				f.Default = []int32(nil)
			}
			flags.Int32SliceVarP(val, f.Long, f.Short, f.Default.([]int32), f.Usage)
		case *[]int64:
			if f.Default == nil {
				f.Default = []int64(nil)
			}
			flags.Int64SliceVarP(val, f.Long, f.Short, f.Default.([]int64), f.Usage)
		case *[]int:
			if f.Default == nil {
				f.Default = []int(nil)
			}
			flags.IntSliceVarP(val, f.Long, f.Short, f.Default.([]int), f.Usage)
		case *[]string:
			if f.Default == nil {
				f.Default = []string(nil)
			}
			flags.StringSliceVarP(val, f.Long, f.Short, f.Default.([]string), f.Usage)
		default:
			panic(fmt.Errorf("unexpected flag value type: %T", val))
		}

		if f.Required {
			err := cobra.MarkFlagRequired(flags, f.Long)
			if err != nil {
				panic(fmt.Errorf("could not mark flag required: %w", err))
			}
		}

		if f.Hidden {
			err := flags.MarkHidden(f.Long)
			if err != nil {
				panic(fmt.Errorf("could not mark flag hidden: %w", err))
			}
		}
	}
}
