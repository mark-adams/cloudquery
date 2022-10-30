// Code generated by codegen; DO NOT EDIT.

package acm

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Certificates() *schema.Table {
	return &schema.Table{
		Name:                "aws_acm_certificates",
		Description:         "https://docs.aws.amazon.com/acm/latest/APIReference/API_CertificateDetail.html",
		Resolver:            fetchAcmCertificates,
		PreResourceResolver: getCertificate,
		Multiplex:           client.ServiceAccountRegionMultiplexer("acm"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertificateArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveCertificateTags,
			},
			{
				Name:     "certificate_authority_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertificateAuthorityArn"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainName"),
			},
			{
				Name:     "domain_validation_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DomainValidationOptions"),
			},
			{
				Name:     "extended_key_usages",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExtendedKeyUsages"),
			},
			{
				Name:     "failure_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FailureReason"),
			},
			{
				Name:     "imported_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ImportedAt"),
			},
			{
				Name:     "in_use_by",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("InUseBy"),
			},
			{
				Name:     "issued_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("IssuedAt"),
			},
			{
				Name:     "issuer",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Issuer"),
			},
			{
				Name:     "key_algorithm",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyAlgorithm"),
			},
			{
				Name:     "key_usages",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("KeyUsages"),
			},
			{
				Name:     "not_after",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("NotAfter"),
			},
			{
				Name:     "not_before",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("NotBefore"),
			},
			{
				Name:     "options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Options"),
			},
			{
				Name:     "renewal_eligibility",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RenewalEligibility"),
			},
			{
				Name:     "renewal_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RenewalSummary"),
			},
			{
				Name:     "revocation_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RevocationReason"),
			},
			{
				Name:     "revoked_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("RevokedAt"),
			},
			{
				Name:     "serial",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Serial"),
			},
			{
				Name:     "signature_algorithm",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SignatureAlgorithm"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "subject",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subject"),
			},
			{
				Name:     "subject_alternative_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SubjectAlternativeNames"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}
