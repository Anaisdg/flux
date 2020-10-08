// DO NOT EDIT: This file is autogenerated via the builtin command.

package statsmodels

import (
	ast "github.com/influxdata/flux/ast"
	parser "github.com/influxdata/flux/internal/parser"
)

var FluxTestPackages = []*ast.Package{&ast.Package{
	BaseNode: ast.BaseNode{
		Errors: nil,
		Loc:    nil,
	},
	Files: []*ast.File{&ast.File{
		BaseNode: ast.BaseNode{
			Errors: nil,
			Loc: &ast.SourceLocation{
				End: ast.Position{
					Column: 105,
					Line:   35,
				},
				File:   "linearreg_test.flux",
				Source: "package statsmodels_test\n\nimport \"testing\"\nimport \"contrib/anaisdg/statsmodels\" \n\n\ninData = \"\n#group,false,false,true,true,false,false,true,true,true,true\n#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string,string\n#default,_result,,,,,,,,,\n,result,table,_start,_stop,_time,_value,_field,_measurement,shelter,type\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:43:45Z,7,young,cats,A,calico\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:45:08Z,5,young,cats,A,calico\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:46:25Z,4,young,cats,A,calico\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:48:38Z,3,young,cats,A,calico\n\"\noutData = \"\n#group,false,false,false,true,true,true,true,false,false,true,false,false,false,false,false,true,false,false,false\n#datatype,string,long,double,string,string,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,double,double,double,double,double,string,double,double,double\n#default,_result,,,,,,,,,,,,,,,,,,\n,result,table,N,_field,_measurement,_start,_stop,_time,errors,shelter,slope,sx,sxx,sxy,sy,type,x,y,y_hat\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:43:45Z,0.0899999999999999,A,-1.3,10,30,41,19,calico,1,7,6.7\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:45:08Z,0.16000000000000028,A,-1.3,10,30,41,19,calico,2,5,5.4\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:46:25Z,0.009999999999999929,A,-1.3,10,30,41,19,calico,3,4,4.1\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:48:38Z,0.04000000000000007,A,-1.3,10,30,41,19,calico,4,3,2.8\n\"\n\n\nt_linearRegression = (table=<-) =>\n\ttable\n\t\t|> range(start: 2020-05-21T21:30:48.901Z, stop: 2020-05-21T21:50:48.9Z)\n\t\t|> statsmodels.linearRegression()\n\ntest _linearRegression = () =>\n({input: testing.loadStorage(csv: inData), want: testing.loadMem(csv: outData), fn: t_linearRegression})",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 2,
						Line:   16,
					},
					File:   "linearreg_test.flux",
					Source: "inData = \"\n#group,false,false,true,true,false,false,true,true,true,true\n#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string,string\n#default,_result,,,,,,,,,\n,result,table,_start,_stop,_time,_value,_field,_measurement,shelter,type\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:43:45Z,7,young,cats,A,calico\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:45:08Z,5,young,cats,A,calico\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:46:25Z,4,young,cats,A,calico\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:48:38Z,3,young,cats,A,calico\n\"",
					Start: ast.Position{
						Column: 1,
						Line:   7,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 7,
							Line:   7,
						},
						File:   "linearreg_test.flux",
						Source: "inData",
						Start: ast.Position{
							Column: 1,
							Line:   7,
						},
					},
				},
				Name: "inData",
			},
			Init: &ast.StringLiteral{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 2,
							Line:   16,
						},
						File:   "linearreg_test.flux",
						Source: "\"\n#group,false,false,true,true,false,false,true,true,true,true\n#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string,string\n#default,_result,,,,,,,,,\n,result,table,_start,_stop,_time,_value,_field,_measurement,shelter,type\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:43:45Z,7,young,cats,A,calico\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:45:08Z,5,young,cats,A,calico\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:46:25Z,4,young,cats,A,calico\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:48:38Z,3,young,cats,A,calico\n\"",
						Start: ast.Position{
							Column: 10,
							Line:   7,
						},
					},
				},
				Value: "\n#group,false,false,true,true,false,false,true,true,true,true\n#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string,string\n#default,_result,,,,,,,,,\n,result,table,_start,_stop,_time,_value,_field,_measurement,shelter,type\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:43:45Z,7,young,cats,A,calico\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:45:08Z,5,young,cats,A,calico\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:46:25Z,4,young,cats,A,calico\n,,0,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:48:38Z,3,young,cats,A,calico\n",
			},
		}, &ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 2,
						Line:   26,
					},
					File:   "linearreg_test.flux",
					Source: "outData = \"\n#group,false,false,false,true,true,true,true,false,false,true,false,false,false,false,false,true,false,false,false\n#datatype,string,long,double,string,string,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,double,double,double,double,double,string,double,double,double\n#default,_result,,,,,,,,,,,,,,,,,,\n,result,table,N,_field,_measurement,_start,_stop,_time,errors,shelter,slope,sx,sxx,sxy,sy,type,x,y,y_hat\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:43:45Z,0.0899999999999999,A,-1.3,10,30,41,19,calico,1,7,6.7\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:45:08Z,0.16000000000000028,A,-1.3,10,30,41,19,calico,2,5,5.4\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:46:25Z,0.009999999999999929,A,-1.3,10,30,41,19,calico,3,4,4.1\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:48:38Z,0.04000000000000007,A,-1.3,10,30,41,19,calico,4,3,2.8\n\"",
					Start: ast.Position{
						Column: 1,
						Line:   17,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 8,
							Line:   17,
						},
						File:   "linearreg_test.flux",
						Source: "outData",
						Start: ast.Position{
							Column: 1,
							Line:   17,
						},
					},
				},
				Name: "outData",
			},
			Init: &ast.StringLiteral{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 2,
							Line:   26,
						},
						File:   "linearreg_test.flux",
						Source: "\"\n#group,false,false,false,true,true,true,true,false,false,true,false,false,false,false,false,true,false,false,false\n#datatype,string,long,double,string,string,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,double,double,double,double,double,string,double,double,double\n#default,_result,,,,,,,,,,,,,,,,,,\n,result,table,N,_field,_measurement,_start,_stop,_time,errors,shelter,slope,sx,sxx,sxy,sy,type,x,y,y_hat\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:43:45Z,0.0899999999999999,A,-1.3,10,30,41,19,calico,1,7,6.7\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:45:08Z,0.16000000000000028,A,-1.3,10,30,41,19,calico,2,5,5.4\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:46:25Z,0.009999999999999929,A,-1.3,10,30,41,19,calico,3,4,4.1\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:48:38Z,0.04000000000000007,A,-1.3,10,30,41,19,calico,4,3,2.8\n\"",
						Start: ast.Position{
							Column: 11,
							Line:   17,
						},
					},
				},
				Value: "\n#group,false,false,false,true,true,true,true,false,false,true,false,false,false,false,false,true,false,false,false\n#datatype,string,long,double,string,string,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,double,double,double,double,double,string,double,double,double\n#default,_result,,,,,,,,,,,,,,,,,,\n,result,table,N,_field,_measurement,_start,_stop,_time,errors,shelter,slope,sx,sxx,sxy,sy,type,x,y,y_hat\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:43:45Z,0.0899999999999999,A,-1.3,10,30,41,19,calico,1,7,6.7\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:45:08Z,0.16000000000000028,A,-1.3,10,30,41,19,calico,2,5,5.4\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:46:25Z,0.009999999999999929,A,-1.3,10,30,41,19,calico,3,4,4.1\n,,0,4,young,cats,2020-05-21T21:30:48.901Z,2020-05-21T21:50:48.9Z,2020-05-21T21:48:38Z,0.04000000000000007,A,-1.3,10,30,41,19,calico,4,3,2.8\n",
			},
		}, &ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 36,
						Line:   32,
					},
					File:   "linearreg_test.flux",
					Source: "t_linearRegression = (table=<-) =>\n\ttable\n\t\t|> range(start: 2020-05-21T21:30:48.901Z, stop: 2020-05-21T21:50:48.9Z)\n\t\t|> statsmodels.linearRegression()",
					Start: ast.Position{
						Column: 1,
						Line:   29,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 19,
							Line:   29,
						},
						File:   "linearreg_test.flux",
						Source: "t_linearRegression",
						Start: ast.Position{
							Column: 1,
							Line:   29,
						},
					},
				},
				Name: "t_linearRegression",
			},
			Init: &ast.FunctionExpression{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 36,
							Line:   32,
						},
						File:   "linearreg_test.flux",
						Source: "(table=<-) =>\n\ttable\n\t\t|> range(start: 2020-05-21T21:30:48.901Z, stop: 2020-05-21T21:50:48.9Z)\n\t\t|> statsmodels.linearRegression()",
						Start: ast.Position{
							Column: 22,
							Line:   29,
						},
					},
				},
				Body: &ast.PipeExpression{
					Argument: &ast.PipeExpression{
						Argument: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 7,
										Line:   30,
									},
									File:   "linearreg_test.flux",
									Source: "table",
									Start: ast.Position{
										Column: 2,
										Line:   30,
									},
								},
							},
							Name: "table",
						},
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 74,
									Line:   31,
								},
								File:   "linearreg_test.flux",
								Source: "table\n\t\t|> range(start: 2020-05-21T21:30:48.901Z, stop: 2020-05-21T21:50:48.9Z)",
								Start: ast.Position{
									Column: 2,
									Line:   30,
								},
							},
						},
						Call: &ast.CallExpression{
							Arguments: []ast.Expression{&ast.ObjectExpression{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 73,
											Line:   31,
										},
										File:   "linearreg_test.flux",
										Source: "start: 2020-05-21T21:30:48.901Z, stop: 2020-05-21T21:50:48.9Z",
										Start: ast.Position{
											Column: 12,
											Line:   31,
										},
									},
								},
								Properties: []*ast.Property{&ast.Property{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 43,
												Line:   31,
											},
											File:   "linearreg_test.flux",
											Source: "start: 2020-05-21T21:30:48.901Z",
											Start: ast.Position{
												Column: 12,
												Line:   31,
											},
										},
									},
									Key: &ast.Identifier{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 17,
													Line:   31,
												},
												File:   "linearreg_test.flux",
												Source: "start",
												Start: ast.Position{
													Column: 12,
													Line:   31,
												},
											},
										},
										Name: "start",
									},
									Value: &ast.DateTimeLiteral{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 43,
													Line:   31,
												},
												File:   "linearreg_test.flux",
												Source: "2020-05-21T21:30:48.901Z",
												Start: ast.Position{
													Column: 19,
													Line:   31,
												},
											},
										},
										Value: parser.MustParseTime("2020-05-21T21:30:48.901Z"),
									},
								}, &ast.Property{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 73,
												Line:   31,
											},
											File:   "linearreg_test.flux",
											Source: "stop: 2020-05-21T21:50:48.9Z",
											Start: ast.Position{
												Column: 45,
												Line:   31,
											},
										},
									},
									Key: &ast.Identifier{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 49,
													Line:   31,
												},
												File:   "linearreg_test.flux",
												Source: "stop",
												Start: ast.Position{
													Column: 45,
													Line:   31,
												},
											},
										},
										Name: "stop",
									},
									Value: &ast.DateTimeLiteral{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 73,
													Line:   31,
												},
												File:   "linearreg_test.flux",
												Source: "2020-05-21T21:50:48.9Z",
												Start: ast.Position{
													Column: 51,
													Line:   31,
												},
											},
										},
										Value: parser.MustParseTime("2020-05-21T21:50:48.9Z"),
									},
								}},
								With: nil,
							}},
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 74,
										Line:   31,
									},
									File:   "linearreg_test.flux",
									Source: "range(start: 2020-05-21T21:30:48.901Z, stop: 2020-05-21T21:50:48.9Z)",
									Start: ast.Position{
										Column: 6,
										Line:   31,
									},
								},
							},
							Callee: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 11,
											Line:   31,
										},
										File:   "linearreg_test.flux",
										Source: "range",
										Start: ast.Position{
											Column: 6,
											Line:   31,
										},
									},
								},
								Name: "range",
							},
						},
					},
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 36,
								Line:   32,
							},
							File:   "linearreg_test.flux",
							Source: "table\n\t\t|> range(start: 2020-05-21T21:30:48.901Z, stop: 2020-05-21T21:50:48.9Z)\n\t\t|> statsmodels.linearRegression()",
							Start: ast.Position{
								Column: 2,
								Line:   30,
							},
						},
					},
					Call: &ast.CallExpression{
						Arguments: nil,
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 36,
									Line:   32,
								},
								File:   "linearreg_test.flux",
								Source: "statsmodels.linearRegression()",
								Start: ast.Position{
									Column: 6,
									Line:   32,
								},
							},
						},
						Callee: &ast.MemberExpression{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 34,
										Line:   32,
									},
									File:   "linearreg_test.flux",
									Source: "statsmodels.linearRegression",
									Start: ast.Position{
										Column: 6,
										Line:   32,
									},
								},
							},
							Object: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 17,
											Line:   32,
										},
										File:   "linearreg_test.flux",
										Source: "statsmodels",
										Start: ast.Position{
											Column: 6,
											Line:   32,
										},
									},
								},
								Name: "statsmodels",
							},
							Property: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 34,
											Line:   32,
										},
										File:   "linearreg_test.flux",
										Source: "linearRegression",
										Start: ast.Position{
											Column: 18,
											Line:   32,
										},
									},
								},
								Name: "linearRegression",
							},
						},
					},
				},
				Params: []*ast.Property{&ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 31,
								Line:   29,
							},
							File:   "linearreg_test.flux",
							Source: "table=<-",
							Start: ast.Position{
								Column: 23,
								Line:   29,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 28,
									Line:   29,
								},
								File:   "linearreg_test.flux",
								Source: "table",
								Start: ast.Position{
									Column: 23,
									Line:   29,
								},
							},
						},
						Name: "table",
					},
					Value: &ast.PipeLiteral{BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 31,
								Line:   29,
							},
							File:   "linearreg_test.flux",
							Source: "<-",
							Start: ast.Position{
								Column: 29,
								Line:   29,
							},
						},
					}},
				}},
			},
		}, &ast.TestStatement{
			Assignment: &ast.VariableAssignment{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 105,
							Line:   35,
						},
						File:   "linearreg_test.flux",
						Source: "_linearRegression = () =>\n({input: testing.loadStorage(csv: inData), want: testing.loadMem(csv: outData), fn: t_linearRegression})",
						Start: ast.Position{
							Column: 6,
							Line:   34,
						},
					},
				},
				ID: &ast.Identifier{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 23,
								Line:   34,
							},
							File:   "linearreg_test.flux",
							Source: "_linearRegression",
							Start: ast.Position{
								Column: 6,
								Line:   34,
							},
						},
					},
					Name: "_linearRegression",
				},
				Init: &ast.FunctionExpression{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 105,
								Line:   35,
							},
							File:   "linearreg_test.flux",
							Source: "() =>\n({input: testing.loadStorage(csv: inData), want: testing.loadMem(csv: outData), fn: t_linearRegression})",
							Start: ast.Position{
								Column: 26,
								Line:   34,
							},
						},
					},
					Body: &ast.ParenExpression{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 105,
									Line:   35,
								},
								File:   "linearreg_test.flux",
								Source: "({input: testing.loadStorage(csv: inData), want: testing.loadMem(csv: outData), fn: t_linearRegression})",
								Start: ast.Position{
									Column: 1,
									Line:   35,
								},
							},
						},
						Expression: &ast.ObjectExpression{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 104,
										Line:   35,
									},
									File:   "linearreg_test.flux",
									Source: "{input: testing.loadStorage(csv: inData), want: testing.loadMem(csv: outData), fn: t_linearRegression}",
									Start: ast.Position{
										Column: 2,
										Line:   35,
									},
								},
							},
							Properties: []*ast.Property{&ast.Property{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 42,
											Line:   35,
										},
										File:   "linearreg_test.flux",
										Source: "input: testing.loadStorage(csv: inData)",
										Start: ast.Position{
											Column: 3,
											Line:   35,
										},
									},
								},
								Key: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 8,
												Line:   35,
											},
											File:   "linearreg_test.flux",
											Source: "input",
											Start: ast.Position{
												Column: 3,
												Line:   35,
											},
										},
									},
									Name: "input",
								},
								Value: &ast.CallExpression{
									Arguments: []ast.Expression{&ast.ObjectExpression{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 41,
													Line:   35,
												},
												File:   "linearreg_test.flux",
												Source: "csv: inData",
												Start: ast.Position{
													Column: 30,
													Line:   35,
												},
											},
										},
										Properties: []*ast.Property{&ast.Property{
											BaseNode: ast.BaseNode{
												Errors: nil,
												Loc: &ast.SourceLocation{
													End: ast.Position{
														Column: 41,
														Line:   35,
													},
													File:   "linearreg_test.flux",
													Source: "csv: inData",
													Start: ast.Position{
														Column: 30,
														Line:   35,
													},
												},
											},
											Key: &ast.Identifier{
												BaseNode: ast.BaseNode{
													Errors: nil,
													Loc: &ast.SourceLocation{
														End: ast.Position{
															Column: 33,
															Line:   35,
														},
														File:   "linearreg_test.flux",
														Source: "csv",
														Start: ast.Position{
															Column: 30,
															Line:   35,
														},
													},
												},
												Name: "csv",
											},
											Value: &ast.Identifier{
												BaseNode: ast.BaseNode{
													Errors: nil,
													Loc: &ast.SourceLocation{
														End: ast.Position{
															Column: 41,
															Line:   35,
														},
														File:   "linearreg_test.flux",
														Source: "inData",
														Start: ast.Position{
															Column: 35,
															Line:   35,
														},
													},
												},
												Name: "inData",
											},
										}},
										With: nil,
									}},
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 42,
												Line:   35,
											},
											File:   "linearreg_test.flux",
											Source: "testing.loadStorage(csv: inData)",
											Start: ast.Position{
												Column: 10,
												Line:   35,
											},
										},
									},
									Callee: &ast.MemberExpression{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 29,
													Line:   35,
												},
												File:   "linearreg_test.flux",
												Source: "testing.loadStorage",
												Start: ast.Position{
													Column: 10,
													Line:   35,
												},
											},
										},
										Object: &ast.Identifier{
											BaseNode: ast.BaseNode{
												Errors: nil,
												Loc: &ast.SourceLocation{
													End: ast.Position{
														Column: 17,
														Line:   35,
													},
													File:   "linearreg_test.flux",
													Source: "testing",
													Start: ast.Position{
														Column: 10,
														Line:   35,
													},
												},
											},
											Name: "testing",
										},
										Property: &ast.Identifier{
											BaseNode: ast.BaseNode{
												Errors: nil,
												Loc: &ast.SourceLocation{
													End: ast.Position{
														Column: 29,
														Line:   35,
													},
													File:   "linearreg_test.flux",
													Source: "loadStorage",
													Start: ast.Position{
														Column: 18,
														Line:   35,
													},
												},
											},
											Name: "loadStorage",
										},
									},
								},
							}, &ast.Property{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 79,
											Line:   35,
										},
										File:   "linearreg_test.flux",
										Source: "want: testing.loadMem(csv: outData)",
										Start: ast.Position{
											Column: 44,
											Line:   35,
										},
									},
								},
								Key: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 48,
												Line:   35,
											},
											File:   "linearreg_test.flux",
											Source: "want",
											Start: ast.Position{
												Column: 44,
												Line:   35,
											},
										},
									},
									Name: "want",
								},
								Value: &ast.CallExpression{
									Arguments: []ast.Expression{&ast.ObjectExpression{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 78,
													Line:   35,
												},
												File:   "linearreg_test.flux",
												Source: "csv: outData",
												Start: ast.Position{
													Column: 66,
													Line:   35,
												},
											},
										},
										Properties: []*ast.Property{&ast.Property{
											BaseNode: ast.BaseNode{
												Errors: nil,
												Loc: &ast.SourceLocation{
													End: ast.Position{
														Column: 78,
														Line:   35,
													},
													File:   "linearreg_test.flux",
													Source: "csv: outData",
													Start: ast.Position{
														Column: 66,
														Line:   35,
													},
												},
											},
											Key: &ast.Identifier{
												BaseNode: ast.BaseNode{
													Errors: nil,
													Loc: &ast.SourceLocation{
														End: ast.Position{
															Column: 69,
															Line:   35,
														},
														File:   "linearreg_test.flux",
														Source: "csv",
														Start: ast.Position{
															Column: 66,
															Line:   35,
														},
													},
												},
												Name: "csv",
											},
											Value: &ast.Identifier{
												BaseNode: ast.BaseNode{
													Errors: nil,
													Loc: &ast.SourceLocation{
														End: ast.Position{
															Column: 78,
															Line:   35,
														},
														File:   "linearreg_test.flux",
														Source: "outData",
														Start: ast.Position{
															Column: 71,
															Line:   35,
														},
													},
												},
												Name: "outData",
											},
										}},
										With: nil,
									}},
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 79,
												Line:   35,
											},
											File:   "linearreg_test.flux",
											Source: "testing.loadMem(csv: outData)",
											Start: ast.Position{
												Column: 50,
												Line:   35,
											},
										},
									},
									Callee: &ast.MemberExpression{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 65,
													Line:   35,
												},
												File:   "linearreg_test.flux",
												Source: "testing.loadMem",
												Start: ast.Position{
													Column: 50,
													Line:   35,
												},
											},
										},
										Object: &ast.Identifier{
											BaseNode: ast.BaseNode{
												Errors: nil,
												Loc: &ast.SourceLocation{
													End: ast.Position{
														Column: 57,
														Line:   35,
													},
													File:   "linearreg_test.flux",
													Source: "testing",
													Start: ast.Position{
														Column: 50,
														Line:   35,
													},
												},
											},
											Name: "testing",
										},
										Property: &ast.Identifier{
											BaseNode: ast.BaseNode{
												Errors: nil,
												Loc: &ast.SourceLocation{
													End: ast.Position{
														Column: 65,
														Line:   35,
													},
													File:   "linearreg_test.flux",
													Source: "loadMem",
													Start: ast.Position{
														Column: 58,
														Line:   35,
													},
												},
											},
											Name: "loadMem",
										},
									},
								},
							}, &ast.Property{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 103,
											Line:   35,
										},
										File:   "linearreg_test.flux",
										Source: "fn: t_linearRegression",
										Start: ast.Position{
											Column: 81,
											Line:   35,
										},
									},
								},
								Key: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 83,
												Line:   35,
											},
											File:   "linearreg_test.flux",
											Source: "fn",
											Start: ast.Position{
												Column: 81,
												Line:   35,
											},
										},
									},
									Name: "fn",
								},
								Value: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 103,
												Line:   35,
											},
											File:   "linearreg_test.flux",
											Source: "t_linearRegression",
											Start: ast.Position{
												Column: 85,
												Line:   35,
											},
										},
									},
									Name: "t_linearRegression",
								},
							}},
							With: nil,
						},
					},
					Params: []*ast.Property{},
				},
			},
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 105,
						Line:   35,
					},
					File:   "linearreg_test.flux",
					Source: "test _linearRegression = () =>\n({input: testing.loadStorage(csv: inData), want: testing.loadMem(csv: outData), fn: t_linearRegression})",
					Start: ast.Position{
						Column: 1,
						Line:   34,
					},
				},
			},
		}},
		Imports: []*ast.ImportDeclaration{&ast.ImportDeclaration{
			As: nil,
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 17,
						Line:   3,
					},
					File:   "linearreg_test.flux",
					Source: "import \"testing\"",
					Start: ast.Position{
						Column: 1,
						Line:   3,
					},
				},
			},
			Path: &ast.StringLiteral{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 17,
							Line:   3,
						},
						File:   "linearreg_test.flux",
						Source: "\"testing\"",
						Start: ast.Position{
							Column: 8,
							Line:   3,
						},
					},
				},
				Value: "testing",
			},
		}, &ast.ImportDeclaration{
			As: nil,
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 37,
						Line:   4,
					},
					File:   "linearreg_test.flux",
					Source: "import \"contrib/anaisdg/statsmodels\"",
					Start: ast.Position{
						Column: 1,
						Line:   4,
					},
				},
			},
			Path: &ast.StringLiteral{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 37,
							Line:   4,
						},
						File:   "linearreg_test.flux",
						Source: "\"contrib/anaisdg/statsmodels\"",
						Start: ast.Position{
							Column: 8,
							Line:   4,
						},
					},
				},
				Value: "contrib/anaisdg/statsmodels",
			},
		}},
		Metadata: "parser-type=rust",
		Name:     "linearreg_test.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 25,
						Line:   1,
					},
					File:   "linearreg_test.flux",
					Source: "package statsmodels_test",
					Start: ast.Position{
						Column: 1,
						Line:   1,
					},
				},
			},
			Name: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 25,
							Line:   1,
						},
						File:   "linearreg_test.flux",
						Source: "statsmodels_test",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "statsmodels_test",
			},
		},
	}},
	Package: "statsmodels_test",
	Path:    "contrib/anaisdg/statsmodels",
}}
