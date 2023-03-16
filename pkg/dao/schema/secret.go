package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/seal-io/seal/pkg/dao/schema/mixin"
	"github.com/seal-io/seal/pkg/dao/types/crypto"
	"github.com/seal-io/seal/pkg/dao/types/oid"
)

type Secret struct {
	ent.Schema
}

func (Secret) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
		mixin.Time{},
	}
}

func (Secret) Indexes() []ent.Index {
	return []ent.Index{
		// NB(thxCode): since null project secret belongs to the organization(beyond any project),
		// single unique constraint index cannot cover null column value,
		// so we leverage conditional indexes to run this case.
		index.Fields("projectID", "name").
			Unique().
			Annotations(entsql.IndexAnnotation{
				Where: "project_id IS NOT NULL",
			}),
		index.Fields("name").
			Unique().
			Annotations(entsql.IndexAnnotation{
				Where: "project_id IS NULL",
			}),
	}
}

func (Secret) Fields() []ent.Field {
	return []ent.Field{
		oid.Field("projectID").
			Comment("ID of the project to which the secret belongs, " +
				"empty means sharing for all projects.").
			Optional().
			Immutable(),
		field.String("name").
			Comment("The name of secret.").
			NotEmpty().
			Immutable(),
		crypto.StringField("value").
			Comment("The value of secret, store in string.").
			NotEmpty().
			Sensitive(),
	}
}

func (Secret) Edges() []ent.Edge {
	return []ent.Edge{
		// project 1-* secrets.
		edge.From("project", Project.Type).
			Ref("secrets").
			Field("projectID").
			Comment("Project to which this secret belongs.").
			Unique().
			Immutable(),
	}
}