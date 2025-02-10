package main

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/mburbidg/gkit-parser/gen"
)

type Visitor struct {
}

func NewVisitor() gen.GQLVisitor {
	return &Visitor{}
}

func (v Visitor) Visit(tree antlr.ParseTree) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitChildren(node antlr.RuleNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGqlProgram(ctx *gen.GqlProgramContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitProgramActivity(ctx *gen.ProgramActivityContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSessionActivity(ctx *gen.SessionActivityContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTransactionActivity(ctx *gen.TransactionActivityContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEndTransactionCommand(ctx *gen.EndTransactionCommandContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSessionSetCommand(ctx *gen.SessionSetCommandContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSessionSetSchemaClause(ctx *gen.SessionSetSchemaClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSessionSetGraphClause(ctx *gen.SessionSetGraphClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSessionSetTimeZoneClause(ctx *gen.SessionSetTimeZoneClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSetTimeZoneValue(ctx *gen.SetTimeZoneValueContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSessionSetParameterClause(ctx *gen.SessionSetParameterClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSessionSetGraphParameterClause(ctx *gen.SessionSetGraphParameterClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSessionSetBindingTableParameterClause(ctx *gen.SessionSetBindingTableParameterClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSessionSetValueParameterClause(ctx *gen.SessionSetValueParameterClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSessionSetParameterName(ctx *gen.SessionSetParameterNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSessionResetCommand(ctx *gen.SessionResetCommandContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSessionResetArguments(ctx *gen.SessionResetArgumentsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSessionCloseCommand(ctx *gen.SessionCloseCommandContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSessionParameterSpecification(ctx *gen.SessionParameterSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitStartTransactionCommand(ctx *gen.StartTransactionCommandContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTransactionCharacteristics(ctx *gen.TransactionCharacteristicsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTransactionMode(ctx *gen.TransactionModeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTransactionAccessMode(ctx *gen.TransactionAccessModeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitRollbackCommand(ctx *gen.RollbackCommandContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCommitCommand(ctx *gen.CommitCommandContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNestedProcedureSpecification(ctx *gen.NestedProcedureSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitProcedureSpecification(ctx *gen.ProcedureSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNestedDataModifyingProcedureSpecification(ctx *gen.NestedDataModifyingProcedureSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNestedQuerySpecification(ctx *gen.NestedQuerySpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitProcedureBody(ctx *gen.ProcedureBodyContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBindingVariableDefinitionBlock(ctx *gen.BindingVariableDefinitionBlockContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBindingVariableDefinition(ctx *gen.BindingVariableDefinitionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitStatementBlock(ctx *gen.StatementBlockContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitStatement(ctx *gen.StatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNextStatement(ctx *gen.NextStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphVariableDefinition(ctx *gen.GraphVariableDefinitionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOptTypedGraphInitializer(ctx *gen.OptTypedGraphInitializerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphInitializer(ctx *gen.GraphInitializerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBindingTableVariableDefinition(ctx *gen.BindingTableVariableDefinitionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOptTypedBindingTableInitializer(ctx *gen.OptTypedBindingTableInitializerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBindingTableInitializer(ctx *gen.BindingTableInitializerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitValueVariableDefinition(ctx *gen.ValueVariableDefinitionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOptTypedValueInitializer(ctx *gen.OptTypedValueInitializerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitValueInitializer(ctx *gen.ValueInitializerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphExpression(ctx *gen.GraphExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCurrentGraph(ctx *gen.CurrentGraphContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBindingTableExpression(ctx *gen.BindingTableExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNestedBindingTableQuerySpecification(ctx *gen.NestedBindingTableQuerySpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitObjectExpressionPrimary(ctx *gen.ObjectExpressionPrimaryContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLinearCatalogModifyingStatement(ctx *gen.LinearCatalogModifyingStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimpleCatalogModifyingStatement(ctx *gen.SimpleCatalogModifyingStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPrimitiveCatalogModifyingStatement(ctx *gen.PrimitiveCatalogModifyingStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCreateSchemaStatement(ctx *gen.CreateSchemaStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDropSchemaStatement(ctx *gen.DropSchemaStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCreateGraphStatement(ctx *gen.CreateGraphStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOpenGraphType(ctx *gen.OpenGraphTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOfGraphType(ctx *gen.OfGraphTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphTypeLikeGraph(ctx *gen.GraphTypeLikeGraphContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphSource(ctx *gen.GraphSourceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDropGraphStatement(ctx *gen.DropGraphStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCreateGraphTypeStatement(ctx *gen.CreateGraphTypeStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphTypeSource(ctx *gen.GraphTypeSourceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCopyOfGraphType(ctx *gen.CopyOfGraphTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDropGraphTypeStatement(ctx *gen.DropGraphTypeStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCallCatalogModifyingProcedureStatement(ctx *gen.CallCatalogModifyingProcedureStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLinearDataModifyingStatement(ctx *gen.LinearDataModifyingStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFocusedLinearDataModifyingStatement(ctx *gen.FocusedLinearDataModifyingStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFocusedLinearDataModifyingStatementBody(ctx *gen.FocusedLinearDataModifyingStatementBodyContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFocusedNestedDataModifyingProcedureSpecification(ctx *gen.FocusedNestedDataModifyingProcedureSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAmbientLinearDataModifyingStatement(ctx *gen.AmbientLinearDataModifyingStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAmbientLinearDataModifyingStatementBody(ctx *gen.AmbientLinearDataModifyingStatementBodyContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimpleLinearDataAccessingStatement(ctx *gen.SimpleLinearDataAccessingStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimpleDataAccessingStatement(ctx *gen.SimpleDataAccessingStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimpleDataModifyingStatement(ctx *gen.SimpleDataModifyingStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPrimitiveDataModifyingStatement(ctx *gen.PrimitiveDataModifyingStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitInsertStatement(ctx *gen.InsertStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSetStatement(ctx *gen.SetStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSetItemList(ctx *gen.SetItemListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSetItem(ctx *gen.SetItemContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSetPropertyItem(ctx *gen.SetPropertyItemContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSetAllPropertiesItem(ctx *gen.SetAllPropertiesItemContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSetLabelItem(ctx *gen.SetLabelItemContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitRemoveStatement(ctx *gen.RemoveStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitRemoveItemList(ctx *gen.RemoveItemListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitRemoveItem(ctx *gen.RemoveItemContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitRemovePropertyItem(ctx *gen.RemovePropertyItemContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitRemoveLabelItem(ctx *gen.RemoveLabelItemContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDeleteStatement(ctx *gen.DeleteStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDeleteItemList(ctx *gen.DeleteItemListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDeleteItem(ctx *gen.DeleteItemContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCallDataModifyingProcedureStatement(ctx *gen.CallDataModifyingProcedureStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCompositeQueryStatement(ctx *gen.CompositeQueryStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCompositeQueryExpression(ctx *gen.CompositeQueryExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitQueryConjunction(ctx *gen.QueryConjunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSetOperator(ctx *gen.SetOperatorContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCompositeQueryPrimary(ctx *gen.CompositeQueryPrimaryContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLinearQueryStatement(ctx *gen.LinearQueryStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFocusedLinearQueryStatement(ctx *gen.FocusedLinearQueryStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFocusedLinearQueryStatementPart(ctx *gen.FocusedLinearQueryStatementPartContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFocusedLinearQueryAndPrimitiveResultStatementPart(ctx *gen.FocusedLinearQueryAndPrimitiveResultStatementPartContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFocusedPrimitiveResultStatement(ctx *gen.FocusedPrimitiveResultStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFocusedNestedQuerySpecification(ctx *gen.FocusedNestedQuerySpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAmbientLinearQueryStatement(ctx *gen.AmbientLinearQueryStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimpleLinearQueryStatement(ctx *gen.SimpleLinearQueryStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimpleQueryStatement(ctx *gen.SimpleQueryStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPrimitiveQueryStatement(ctx *gen.PrimitiveQueryStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitMatchStatement(ctx *gen.MatchStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimpleMatchStatement(ctx *gen.SimpleMatchStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOptionalMatchStatement(ctx *gen.OptionalMatchStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOptionalOperand(ctx *gen.OptionalOperandContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitMatchStatementBlock(ctx *gen.MatchStatementBlockContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCallQueryStatement(ctx *gen.CallQueryStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFilterStatement(ctx *gen.FilterStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLetStatement(ctx *gen.LetStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLetVariableDefinitionList(ctx *gen.LetVariableDefinitionListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLetVariableDefinition(ctx *gen.LetVariableDefinitionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitForStatement(ctx *gen.ForStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitForItem(ctx *gen.ForItemContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitForItemAlias(ctx *gen.ForItemAliasContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitForItemSource(ctx *gen.ForItemSourceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitForOrdinalityOrOffset(ctx *gen.ForOrdinalityOrOffsetContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOrderByAndPageStatement(ctx *gen.OrderByAndPageStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPrimitiveResultStatement(ctx *gen.PrimitiveResultStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitReturnStatement(ctx *gen.ReturnStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitReturnStatementBody(ctx *gen.ReturnStatementBodyContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitReturnItemList(ctx *gen.ReturnItemListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitReturnItem(ctx *gen.ReturnItemContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitReturnItemAlias(ctx *gen.ReturnItemAliasContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSelectStatement(ctx *gen.SelectStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSelectItemList(ctx *gen.SelectItemListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSelectItem(ctx *gen.SelectItemContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSelectItemAlias(ctx *gen.SelectItemAliasContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitHavingClause(ctx *gen.HavingClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSelectStatementBody(ctx *gen.SelectStatementBodyContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSelectGraphMatchList(ctx *gen.SelectGraphMatchListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSelectGraphMatch(ctx *gen.SelectGraphMatchContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSelectQuerySpecification(ctx *gen.SelectQuerySpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCallProcedureStatement(ctx *gen.CallProcedureStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitProcedureCall(ctx *gen.ProcedureCallContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitInlineProcedureCall(ctx *gen.InlineProcedureCallContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitVariableScopeClause(ctx *gen.VariableScopeClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBindingVariableReferenceList(ctx *gen.BindingVariableReferenceListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNamedProcedureCall(ctx *gen.NamedProcedureCallContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitProcedureArgumentList(ctx *gen.ProcedureArgumentListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitProcedureArgument(ctx *gen.ProcedureArgumentContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAtSchemaClause(ctx *gen.AtSchemaClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitUseGraphClause(ctx *gen.UseGraphClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphPatternBindingTable(ctx *gen.GraphPatternBindingTableContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphPatternYieldClause(ctx *gen.GraphPatternYieldClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphPatternYieldItemList(ctx *gen.GraphPatternYieldItemListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphPatternYieldItem(ctx *gen.GraphPatternYieldItemContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphPattern(ctx *gen.GraphPatternContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitMatchMode(ctx *gen.MatchModeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitRepeatableElementsMatchMode(ctx *gen.RepeatableElementsMatchModeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDifferentEdgesMatchMode(ctx *gen.DifferentEdgesMatchModeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitElementBindingsOrElements(ctx *gen.ElementBindingsOrElementsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeBindingsOrEdges(ctx *gen.EdgeBindingsOrEdgesContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathPatternList(ctx *gen.PathPatternListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathPattern(ctx *gen.PathPatternContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathVariableDeclaration(ctx *gen.PathVariableDeclarationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitKeepClause(ctx *gen.KeepClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphPatternWhereClause(ctx *gen.GraphPatternWhereClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitInsertGraphPattern(ctx *gen.InsertGraphPatternContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitInsertPathPatternList(ctx *gen.InsertPathPatternListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitInsertPathPattern(ctx *gen.InsertPathPatternContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitInsertNodePattern(ctx *gen.InsertNodePatternContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitInsertEdgePattern(ctx *gen.InsertEdgePatternContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitInsertEdgePointingLeft(ctx *gen.InsertEdgePointingLeftContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitInsertEdgePointingRight(ctx *gen.InsertEdgePointingRightContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitInsertEdgeUndirected(ctx *gen.InsertEdgeUndirectedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitInsertElementPatternFiller(ctx *gen.InsertElementPatternFillerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLabelAndPropertySetSpecification(ctx *gen.LabelAndPropertySetSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathPatternPrefix(ctx *gen.PathPatternPrefixContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathModePrefix(ctx *gen.PathModePrefixContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathMode(ctx *gen.PathModeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathSearchPrefix(ctx *gen.PathSearchPrefixContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAllPathSearch(ctx *gen.AllPathSearchContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathOrPaths(ctx *gen.PathOrPathsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAnyPathSearch(ctx *gen.AnyPathSearchContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNumberOfPaths(ctx *gen.NumberOfPathsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitShortestPathSearch(ctx *gen.ShortestPathSearchContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAllShortestPathSearch(ctx *gen.AllShortestPathSearchContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAnyShortestPathSearch(ctx *gen.AnyShortestPathSearchContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCountedShortestPathSearch(ctx *gen.CountedShortestPathSearchContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCountedShortestGroupSearch(ctx *gen.CountedShortestGroupSearchContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNumberOfGroups(ctx *gen.NumberOfGroupsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPpePathTerm(ctx *gen.PpePathTermContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPpeMultisetAlternation(ctx *gen.PpeMultisetAlternationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPpePatternUnion(ctx *gen.PpePatternUnionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathTerm(ctx *gen.PathTermContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPfPathPrimary(ctx *gen.PfPathPrimaryContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPfQuantifiedPathPrimary(ctx *gen.PfQuantifiedPathPrimaryContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPfQuestionedPathPrimary(ctx *gen.PfQuestionedPathPrimaryContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPpElementPattern(ctx *gen.PpElementPatternContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPpParenthesizedPathPatternExpression(ctx *gen.PpParenthesizedPathPatternExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPpSimplifiedPathPatternExpression(ctx *gen.PpSimplifiedPathPatternExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitElementPattern(ctx *gen.ElementPatternContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNodePattern(ctx *gen.NodePatternContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitElementPatternFiller(ctx *gen.ElementPatternFillerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitElementVariableDeclaration(ctx *gen.ElementVariableDeclarationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitIsLabelExpression(ctx *gen.IsLabelExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitIsOrColon(ctx *gen.IsOrColonContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitElementPatternPredicate(ctx *gen.ElementPatternPredicateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitElementPatternWhereClause(ctx *gen.ElementPatternWhereClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitElementPropertySpecification(ctx *gen.ElementPropertySpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPropertyKeyValuePairList(ctx *gen.PropertyKeyValuePairListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPropertyKeyValuePair(ctx *gen.PropertyKeyValuePairContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgePattern(ctx *gen.EdgePatternContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFullEdgePattern(ctx *gen.FullEdgePatternContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFullEdgePointingLeft(ctx *gen.FullEdgePointingLeftContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFullEdgeUndirected(ctx *gen.FullEdgeUndirectedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFullEdgePointingRight(ctx *gen.FullEdgePointingRightContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFullEdgeLeftOrUndirected(ctx *gen.FullEdgeLeftOrUndirectedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFullEdgeUndirectedOrRight(ctx *gen.FullEdgeUndirectedOrRightContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFullEdgeLeftOrRight(ctx *gen.FullEdgeLeftOrRightContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFullEdgeAnyDirection(ctx *gen.FullEdgeAnyDirectionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAbbreviatedEdgePattern(ctx *gen.AbbreviatedEdgePatternContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitParenthesizedPathPatternExpression(ctx *gen.ParenthesizedPathPatternExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSubpathVariableDeclaration(ctx *gen.SubpathVariableDeclarationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitParenthesizedPathPatternWhereClause(ctx *gen.ParenthesizedPathPatternWhereClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLabelExpressionNegation(ctx *gen.LabelExpressionNegationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLabelExpressionDisjunction(ctx *gen.LabelExpressionDisjunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLabelExpressionParenthesized(ctx *gen.LabelExpressionParenthesizedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLabelExpressionWildcard(ctx *gen.LabelExpressionWildcardContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLabelExpressionConjunction(ctx *gen.LabelExpressionConjunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLabelExpressionName(ctx *gen.LabelExpressionNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathVariableReference(ctx *gen.PathVariableReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitElementVariableReference(ctx *gen.ElementVariableReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphPatternQuantifier(ctx *gen.GraphPatternQuantifierContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFixedQuantifier(ctx *gen.FixedQuantifierContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGeneralQuantifier(ctx *gen.GeneralQuantifierContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLowerBound(ctx *gen.LowerBoundContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitUpperBound(ctx *gen.UpperBoundContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedPathPatternExpression(ctx *gen.SimplifiedPathPatternExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedDefaultingLeft(ctx *gen.SimplifiedDefaultingLeftContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedDefaultingUndirected(ctx *gen.SimplifiedDefaultingUndirectedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedDefaultingRight(ctx *gen.SimplifiedDefaultingRightContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedDefaultingLeftOrUndirected(ctx *gen.SimplifiedDefaultingLeftOrUndirectedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedDefaultingUndirectedOrRight(ctx *gen.SimplifiedDefaultingUndirectedOrRightContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedDefaultingLeftOrRight(ctx *gen.SimplifiedDefaultingLeftOrRightContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedDefaultingAnyDirection(ctx *gen.SimplifiedDefaultingAnyDirectionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedContents(ctx *gen.SimplifiedContentsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedPathUnion(ctx *gen.SimplifiedPathUnionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedMultisetAlternation(ctx *gen.SimplifiedMultisetAlternationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedFactorLowLabel(ctx *gen.SimplifiedFactorLowLabelContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedConcatenationLabel(ctx *gen.SimplifiedConcatenationLabelContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedConjunctionLabel(ctx *gen.SimplifiedConjunctionLabelContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedFactorHighLabel(ctx *gen.SimplifiedFactorHighLabelContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedFactorHigh(ctx *gen.SimplifiedFactorHighContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedQuantified(ctx *gen.SimplifiedQuantifiedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedQuestioned(ctx *gen.SimplifiedQuestionedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedTertiary(ctx *gen.SimplifiedTertiaryContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedDirectionOverride(ctx *gen.SimplifiedDirectionOverrideContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedOverrideLeft(ctx *gen.SimplifiedOverrideLeftContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedOverrideUndirected(ctx *gen.SimplifiedOverrideUndirectedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedOverrideRight(ctx *gen.SimplifiedOverrideRightContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedOverrideLeftOrUndirected(ctx *gen.SimplifiedOverrideLeftOrUndirectedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedOverrideUndirectedOrRight(ctx *gen.SimplifiedOverrideUndirectedOrRightContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedOverrideLeftOrRight(ctx *gen.SimplifiedOverrideLeftOrRightContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedOverrideAnyDirection(ctx *gen.SimplifiedOverrideAnyDirectionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedSecondary(ctx *gen.SimplifiedSecondaryContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedNegation(ctx *gen.SimplifiedNegationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimplifiedPrimary(ctx *gen.SimplifiedPrimaryContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitWhereClause(ctx *gen.WhereClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitYieldClause(ctx *gen.YieldClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitYieldItemList(ctx *gen.YieldItemListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitYieldItem(ctx *gen.YieldItemContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitYieldItemName(ctx *gen.YieldItemNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitYieldItemAlias(ctx *gen.YieldItemAliasContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGroupByClause(ctx *gen.GroupByClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGroupingElementList(ctx *gen.GroupingElementListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGroupingElement(ctx *gen.GroupingElementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEmptyGroupingSet(ctx *gen.EmptyGroupingSetContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOrderByClause(ctx *gen.OrderByClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSortSpecificationList(ctx *gen.SortSpecificationListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSortSpecification(ctx *gen.SortSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSortKey(ctx *gen.SortKeyContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOrderingSpecification(ctx *gen.OrderingSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNullOrdering(ctx *gen.NullOrderingContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLimitClause(ctx *gen.LimitClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOffsetClause(ctx *gen.OffsetClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOffsetSynonym(ctx *gen.OffsetSynonymContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSchemaReference(ctx *gen.SchemaReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAbsoluteCatalogSchemaReference(ctx *gen.AbsoluteCatalogSchemaReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCatalogSchemaParentAndName(ctx *gen.CatalogSchemaParentAndNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitRelativeCatalogSchemaReference(ctx *gen.RelativeCatalogSchemaReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPredefinedSchemaReference(ctx *gen.PredefinedSchemaReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAbsoluteDirectoryPath(ctx *gen.AbsoluteDirectoryPathContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitRelativeDirectoryPath(ctx *gen.RelativeDirectoryPathContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimpleDirectoryPath(ctx *gen.SimpleDirectoryPathContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphReference(ctx *gen.GraphReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCatalogGraphParentAndName(ctx *gen.CatalogGraphParentAndNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitHomeGraph(ctx *gen.HomeGraphContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphTypeReference(ctx *gen.GraphTypeReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCatalogGraphTypeParentAndName(ctx *gen.CatalogGraphTypeParentAndNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBindingTableReference(ctx *gen.BindingTableReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitProcedureReference(ctx *gen.ProcedureReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCatalogProcedureParentAndName(ctx *gen.CatalogProcedureParentAndNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCatalogObjectParentReference(ctx *gen.CatalogObjectParentReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitReferenceParameterSpecification(ctx *gen.ReferenceParameterSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNestedGraphTypeSpecification(ctx *gen.NestedGraphTypeSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphTypeSpecificationBody(ctx *gen.GraphTypeSpecificationBodyContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitElementTypeList(ctx *gen.ElementTypeListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitElementTypeSpecification(ctx *gen.ElementTypeSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNodeTypeSpecification(ctx *gen.NodeTypeSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNodeTypePattern(ctx *gen.NodeTypePatternContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNodeTypePhrase(ctx *gen.NodeTypePhraseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNodeTypePhraseFiller(ctx *gen.NodeTypePhraseFillerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNodeTypeFiller(ctx *gen.NodeTypeFillerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLocalNodeTypeAlias(ctx *gen.LocalNodeTypeAliasContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNodeTypeImpliedContent(ctx *gen.NodeTypeImpliedContentContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNodeTypeKeyLabelSet(ctx *gen.NodeTypeKeyLabelSetContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNodeTypeLabelSet(ctx *gen.NodeTypeLabelSetContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNodeTypePropertyTypes(ctx *gen.NodeTypePropertyTypesContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeTypeSpecification(ctx *gen.EdgeTypeSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeTypePattern(ctx *gen.EdgeTypePatternContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeTypePhrase(ctx *gen.EdgeTypePhraseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeTypePhraseFiller(ctx *gen.EdgeTypePhraseFillerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeTypeFiller(ctx *gen.EdgeTypeFillerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeTypeImpliedContent(ctx *gen.EdgeTypeImpliedContentContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeTypeKeyLabelSet(ctx *gen.EdgeTypeKeyLabelSetContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeTypeLabelSet(ctx *gen.EdgeTypeLabelSetContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeTypePropertyTypes(ctx *gen.EdgeTypePropertyTypesContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeTypePatternDirected(ctx *gen.EdgeTypePatternDirectedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeTypePatternPointingRight(ctx *gen.EdgeTypePatternPointingRightContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeTypePatternPointingLeft(ctx *gen.EdgeTypePatternPointingLeftContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeTypePatternUndirected(ctx *gen.EdgeTypePatternUndirectedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitArcTypePointingRight(ctx *gen.ArcTypePointingRightContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitArcTypePointingLeft(ctx *gen.ArcTypePointingLeftContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitArcTypeUndirected(ctx *gen.ArcTypeUndirectedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSourceNodeTypeReference(ctx *gen.SourceNodeTypeReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDestinationNodeTypeReference(ctx *gen.DestinationNodeTypeReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeKind(ctx *gen.EdgeKindContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEndpointPairPhrase(ctx *gen.EndpointPairPhraseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEndpointPair(ctx *gen.EndpointPairContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEndpointPairDirected(ctx *gen.EndpointPairDirectedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEndpointPairPointingRight(ctx *gen.EndpointPairPointingRightContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEndpointPairPointingLeft(ctx *gen.EndpointPairPointingLeftContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEndpointPairUndirected(ctx *gen.EndpointPairUndirectedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitConnectorPointingRight(ctx *gen.ConnectorPointingRightContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitConnectorUndirected(ctx *gen.ConnectorUndirectedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSourceNodeTypeAlias(ctx *gen.SourceNodeTypeAliasContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDestinationNodeTypeAlias(ctx *gen.DestinationNodeTypeAliasContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLabelSetPhrase(ctx *gen.LabelSetPhraseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLabelSetSpecification(ctx *gen.LabelSetSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPropertyTypesSpecification(ctx *gen.PropertyTypesSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPropertyTypeList(ctx *gen.PropertyTypeListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPropertyType(ctx *gen.PropertyTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPropertyValueType(ctx *gen.PropertyValueTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBindingTableType(ctx *gen.BindingTableTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDynamicPropertyValueTypeLabel(ctx *gen.DynamicPropertyValueTypeLabelContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitClosedDynamicUnionTypeAtl1(ctx *gen.ClosedDynamicUnionTypeAtl1Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitClosedDynamicUnionTypeAtl2(ctx *gen.ClosedDynamicUnionTypeAtl2Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathValueTypeLabel(ctx *gen.PathValueTypeLabelContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitListValueTypeAlt3(ctx *gen.ListValueTypeAlt3Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitListValueTypeAlt2(ctx *gen.ListValueTypeAlt2Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitListValueTypeAlt1(ctx *gen.ListValueTypeAlt1Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPredefinedTypeLabel(ctx *gen.PredefinedTypeLabelContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitRecordTypeLabel(ctx *gen.RecordTypeLabelContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOpenDynamicUnionTypeLabel(ctx *gen.OpenDynamicUnionTypeLabelContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTyped(ctx *gen.TypedContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPredefinedType(ctx *gen.PredefinedTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBooleanType(ctx *gen.BooleanTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCharacterStringType(ctx *gen.CharacterStringTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitByteStringType(ctx *gen.ByteStringTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitMinLength(ctx *gen.MinLengthContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitMaxLength(ctx *gen.MaxLengthContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFixedLength(ctx *gen.FixedLengthContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNumericType(ctx *gen.NumericTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitExactNumericType(ctx *gen.ExactNumericTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBinaryExactNumericType(ctx *gen.BinaryExactNumericTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSignedBinaryExactNumericType(ctx *gen.SignedBinaryExactNumericTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitUnsignedBinaryExactNumericType(ctx *gen.UnsignedBinaryExactNumericTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitVerboseBinaryExactNumericType(ctx *gen.VerboseBinaryExactNumericTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDecimalExactNumericType(ctx *gen.DecimalExactNumericTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPrecision(ctx *gen.PrecisionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitScale(ctx *gen.ScaleContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitApproximateNumericType(ctx *gen.ApproximateNumericTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTemporalType(ctx *gen.TemporalTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTemporalInstantType(ctx *gen.TemporalInstantTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDatetimeType(ctx *gen.DatetimeTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLocaldatetimeType(ctx *gen.LocaldatetimeTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDateType(ctx *gen.DateTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTimeType(ctx *gen.TimeTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLocaltimeType(ctx *gen.LocaltimeTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTemporalDurationType(ctx *gen.TemporalDurationTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTemporalDurationQualifier(ctx *gen.TemporalDurationQualifierContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitReferenceValueType(ctx *gen.ReferenceValueTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitImmaterialValueType(ctx *gen.ImmaterialValueTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNullType(ctx *gen.NullTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEmptyType(ctx *gen.EmptyTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphReferenceValueType(ctx *gen.GraphReferenceValueTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitClosedGraphReferenceValueType(ctx *gen.ClosedGraphReferenceValueTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOpenGraphReferenceValueType(ctx *gen.OpenGraphReferenceValueTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBindingTableReferenceValueType(ctx *gen.BindingTableReferenceValueTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNodeReferenceValueType(ctx *gen.NodeReferenceValueTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitClosedNodeReferenceValueType(ctx *gen.ClosedNodeReferenceValueTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOpenNodeReferenceValueType(ctx *gen.OpenNodeReferenceValueTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeReferenceValueType(ctx *gen.EdgeReferenceValueTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitClosedEdgeReferenceValueType(ctx *gen.ClosedEdgeReferenceValueTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitOpenEdgeReferenceValueType(ctx *gen.OpenEdgeReferenceValueTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathValueType(ctx *gen.PathValueTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitListValueTypeName(ctx *gen.ListValueTypeNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitListValueTypeNameSynonym(ctx *gen.ListValueTypeNameSynonymContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitRecordType(ctx *gen.RecordTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFieldTypesSpecification(ctx *gen.FieldTypesSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFieldTypeList(ctx *gen.FieldTypeListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNotNull(ctx *gen.NotNullContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFieldType(ctx *gen.FieldTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSearchCondition(ctx *gen.SearchConditionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPredicate(ctx *gen.PredicateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCompOp(ctx *gen.CompOpContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitExistsPredicate(ctx *gen.ExistsPredicateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNullPredicate(ctx *gen.NullPredicateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNullPredicatePart2(ctx *gen.NullPredicatePart2Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitValueTypePredicate(ctx *gen.ValueTypePredicateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitValueTypePredicatePart2(ctx *gen.ValueTypePredicatePart2Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNormalizedPredicatePart2(ctx *gen.NormalizedPredicatePart2Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDirectedPredicate(ctx *gen.DirectedPredicateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDirectedPredicatePart2(ctx *gen.DirectedPredicatePart2Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLabeledPredicate(ctx *gen.LabeledPredicateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLabeledPredicatePart2(ctx *gen.LabeledPredicatePart2Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitIsLabeledOrColon(ctx *gen.IsLabeledOrColonContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSourceDestinationPredicate(ctx *gen.SourceDestinationPredicateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNodeReference(ctx *gen.NodeReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSourcePredicatePart2(ctx *gen.SourcePredicatePart2Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDestinationPredicatePart2(ctx *gen.DestinationPredicatePart2Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeReference(ctx *gen.EdgeReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAll_differentPredicate(ctx *gen.All_differentPredicateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSamePredicate(ctx *gen.SamePredicateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitProperty_existsPredicate(ctx *gen.Property_existsPredicateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitConjunctiveExprAlt(ctx *gen.ConjunctiveExprAltContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPropertyGraphExprAlt(ctx *gen.PropertyGraphExprAltContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitMultDivExprAlt(ctx *gen.MultDivExprAltContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBindingTableExprAlt(ctx *gen.BindingTableExprAltContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSignedExprAlt(ctx *gen.SignedExprAltContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitIsNotExprAlt(ctx *gen.IsNotExprAltContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNormalizedPredicateExprAlt(ctx *gen.NormalizedPredicateExprAltContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNotExprAlt(ctx *gen.NotExprAltContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitValueFunctionExprAlt(ctx *gen.ValueFunctionExprAltContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitConcatenationExprAlt(ctx *gen.ConcatenationExprAltContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDisjunctiveExprAlt(ctx *gen.DisjunctiveExprAltContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitComparisonExprAlt(ctx *gen.ComparisonExprAltContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPrimaryExprAlt(ctx *gen.PrimaryExprAltContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAddSubtractExprAlt(ctx *gen.AddSubtractExprAltContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPredicateExprAlt(ctx *gen.PredicateExprAltContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitValueFunction(ctx *gen.ValueFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBooleanValueExpression(ctx *gen.BooleanValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCharacterOrByteStringFunction(ctx *gen.CharacterOrByteStringFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSubCharacterOrByteString(ctx *gen.SubCharacterOrByteStringContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTrimSingleCharacterOrByteString(ctx *gen.TrimSingleCharacterOrByteStringContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFoldCharacterString(ctx *gen.FoldCharacterStringContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTrimMultiCharacterCharacterString(ctx *gen.TrimMultiCharacterCharacterStringContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNormalizeCharacterString(ctx *gen.NormalizeCharacterStringContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNodeReferenceValueExpression(ctx *gen.NodeReferenceValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeReferenceValueExpression(ctx *gen.EdgeReferenceValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAggregatingValueExpression(ctx *gen.AggregatingValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitValueExpressionPrimary(ctx *gen.ValueExpressionPrimaryContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitParenthesizedValueExpression(ctx *gen.ParenthesizedValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNonParenthesizedValueExpressionPrimary(ctx *gen.NonParenthesizedValueExpressionPrimaryContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNonParenthesizedValueExpressionPrimarySpecialCase(ctx *gen.NonParenthesizedValueExpressionPrimarySpecialCaseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitUnsignedValueSpecification(ctx *gen.UnsignedValueSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNonNegativeIntegerSpecification(ctx *gen.NonNegativeIntegerSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGeneralValueSpecification(ctx *gen.GeneralValueSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDynamicParameterSpecification(ctx *gen.DynamicParameterSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLetValueExpression(ctx *gen.LetValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitValueQueryExpression(ctx *gen.ValueQueryExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCaseExpression(ctx *gen.CaseExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCaseAbbreviation(ctx *gen.CaseAbbreviationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCaseSpecification(ctx *gen.CaseSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimpleCase(ctx *gen.SimpleCaseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSearchedCase(ctx *gen.SearchedCaseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSimpleWhenClause(ctx *gen.SimpleWhenClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSearchedWhenClause(ctx *gen.SearchedWhenClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitElseClause(ctx *gen.ElseClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCaseOperand(ctx *gen.CaseOperandContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitWhenOperandList(ctx *gen.WhenOperandListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitWhenOperand(ctx *gen.WhenOperandContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitResult(ctx *gen.ResultContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitResultExpression(ctx *gen.ResultExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCastSpecification(ctx *gen.CastSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCastOperand(ctx *gen.CastOperandContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCastTarget(ctx *gen.CastTargetContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAggregateFunction(ctx *gen.AggregateFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGeneralSetFunction(ctx *gen.GeneralSetFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBinarySetFunction(ctx *gen.BinarySetFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGeneralSetFunctionType(ctx *gen.GeneralSetFunctionTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSetQuantifier(ctx *gen.SetQuantifierContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBinarySetFunctionType(ctx *gen.BinarySetFunctionTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDependentValueExpression(ctx *gen.DependentValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitIndependentValueExpression(ctx *gen.IndependentValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitElement_idFunction(ctx *gen.Element_idFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBindingVariableReference(ctx *gen.BindingVariableReferenceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathValueExpression(ctx *gen.PathValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathValueConstructor(ctx *gen.PathValueConstructorContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathValueConstructorByEnumeration(ctx *gen.PathValueConstructorByEnumerationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathElementList(ctx *gen.PathElementListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathElementListStart(ctx *gen.PathElementListStartContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathElementListStep(ctx *gen.PathElementListStepContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitListValueExpression(ctx *gen.ListValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitListValueFunction(ctx *gen.ListValueFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTrimListFunction(ctx *gen.TrimListFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitElementsFunction(ctx *gen.ElementsFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitListValueConstructor(ctx *gen.ListValueConstructorContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitListValueConstructorByEnumeration(ctx *gen.ListValueConstructorByEnumerationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitListElementList(ctx *gen.ListElementListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitListElement(ctx *gen.ListElementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitRecordConstructor(ctx *gen.RecordConstructorContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFieldsSpecification(ctx *gen.FieldsSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFieldList(ctx *gen.FieldListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitField(ctx *gen.FieldContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTruthValue(ctx *gen.TruthValueContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNumericValueExpression(ctx *gen.NumericValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNumericValueFunction(ctx *gen.NumericValueFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLengthExpression(ctx *gen.LengthExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCardinalityExpression(ctx *gen.CardinalityExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCardinalityExpressionArgument(ctx *gen.CardinalityExpressionArgumentContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCharLengthExpression(ctx *gen.CharLengthExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitByteLengthExpression(ctx *gen.ByteLengthExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathLengthExpression(ctx *gen.PathLengthExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitAbsoluteValueExpression(ctx *gen.AbsoluteValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitModulusExpression(ctx *gen.ModulusExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNumericValueExpressionDividend(ctx *gen.NumericValueExpressionDividendContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNumericValueExpressionDivisor(ctx *gen.NumericValueExpressionDivisorContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTrigonometricFunction(ctx *gen.TrigonometricFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTrigonometricFunctionName(ctx *gen.TrigonometricFunctionNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGeneralLogarithmFunction(ctx *gen.GeneralLogarithmFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGeneralLogarithmBase(ctx *gen.GeneralLogarithmBaseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGeneralLogarithmArgument(ctx *gen.GeneralLogarithmArgumentContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCommonLogarithm(ctx *gen.CommonLogarithmContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNaturalLogarithm(ctx *gen.NaturalLogarithmContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitExponentialFunction(ctx *gen.ExponentialFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPowerFunction(ctx *gen.PowerFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNumericValueExpressionBase(ctx *gen.NumericValueExpressionBaseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNumericValueExpressionExponent(ctx *gen.NumericValueExpressionExponentContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSquareRoot(ctx *gen.SquareRootContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFloorFunction(ctx *gen.FloorFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCeilingFunction(ctx *gen.CeilingFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCharacterStringValueExpression(ctx *gen.CharacterStringValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitByteStringValueExpression(ctx *gen.ByteStringValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTrimOperands(ctx *gen.TrimOperandsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTrimCharacterOrByteStringSource(ctx *gen.TrimCharacterOrByteStringSourceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTrimSpecification(ctx *gen.TrimSpecificationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTrimCharacterOrByteString(ctx *gen.TrimCharacterOrByteStringContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNormalForm(ctx *gen.NormalFormContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitStringLength(ctx *gen.StringLengthContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDatetimeValueExpression(ctx *gen.DatetimeValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDatetimeValueFunction(ctx *gen.DatetimeValueFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDateFunction(ctx *gen.DateFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTimeFunction(ctx *gen.TimeFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLocaltimeFunction(ctx *gen.LocaltimeFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDatetimeFunction(ctx *gen.DatetimeFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLocaldatetimeFunction(ctx *gen.LocaldatetimeFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDateFunctionParameters(ctx *gen.DateFunctionParametersContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTimeFunctionParameters(ctx *gen.TimeFunctionParametersContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDatetimeFunctionParameters(ctx *gen.DatetimeFunctionParametersContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDurationValueExpression(ctx *gen.DurationValueExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDatetimeSubtraction(ctx *gen.DatetimeSubtractionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDatetimeSubtractionParameters(ctx *gen.DatetimeSubtractionParametersContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDatetimeValueExpression1(ctx *gen.DatetimeValueExpression1Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDatetimeValueExpression2(ctx *gen.DatetimeValueExpression2Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDurationValueFunction(ctx *gen.DurationValueFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDurationFunction(ctx *gen.DurationFunctionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDurationFunctionParameters(ctx *gen.DurationFunctionParametersContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitObjectName(ctx *gen.ObjectNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitObjectNameOrBindingVariable(ctx *gen.ObjectNameOrBindingVariableContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDirectoryName(ctx *gen.DirectoryNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSchemaName(ctx *gen.SchemaNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphName(ctx *gen.GraphNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDelimitedGraphName(ctx *gen.DelimitedGraphNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGraphTypeName(ctx *gen.GraphTypeNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNodeTypeName(ctx *gen.NodeTypeNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeTypeName(ctx *gen.EdgeTypeNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBindingTableName(ctx *gen.BindingTableNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDelimitedBindingTableName(ctx *gen.DelimitedBindingTableNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitProcedureName(ctx *gen.ProcedureNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitLabelName(ctx *gen.LabelNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPropertyName(ctx *gen.PropertyNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitFieldName(ctx *gen.FieldNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitElementVariable(ctx *gen.ElementVariableContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitPathVariable(ctx *gen.PathVariableContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitSubpathVariable(ctx *gen.SubpathVariableContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitBindingVariable(ctx *gen.BindingVariableContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitUnsignedLiteral(ctx *gen.UnsignedLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitGeneralLiteral(ctx *gen.GeneralLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTemporalLiteral(ctx *gen.TemporalLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDateLiteral(ctx *gen.DateLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTimeLiteral(ctx *gen.TimeLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDatetimeLiteral(ctx *gen.DatetimeLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitListLiteral(ctx *gen.ListLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitRecordLiteral(ctx *gen.RecordLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitIdentifier(ctx *gen.IdentifierContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitRegularIdentifier(ctx *gen.RegularIdentifierContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTimeZoneString(ctx *gen.TimeZoneStringContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitCharacterStringLiteral(ctx *gen.CharacterStringLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitUnsignedNumericLiteral(ctx *gen.UnsignedNumericLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitExactNumericLiteral(ctx *gen.ExactNumericLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitApproximateNumericLiteral(ctx *gen.ApproximateNumericLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitUnsignedInteger(ctx *gen.UnsignedIntegerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitUnsignedDecimalInteger(ctx *gen.UnsignedDecimalIntegerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNullLiteral(ctx *gen.NullLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDateString(ctx *gen.DateStringContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitTimeString(ctx *gen.TimeStringContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDatetimeString(ctx *gen.DatetimeStringContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDurationLiteral(ctx *gen.DurationLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitDurationString(ctx *gen.DurationStringContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNodeSynonym(ctx *gen.NodeSynonymContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgesSynonym(ctx *gen.EdgesSynonymContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitEdgeSynonym(ctx *gen.EdgeSynonymContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v Visitor) VisitNonReservedWords(ctx *gen.NonReservedWordsContext) interface{} {
	//TODO implement me
	panic("implement me")
}
