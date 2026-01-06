
package converting

import . "TypeGo/core"



func PrintFor(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	convertData.SetErrorFunction("PrintFor")
	
	PrintTokensNoNL(convertData, blockData, nestCount); 
	if convertData.IsError() {
		return; 
	}
	convertData.AppendChar(' '); 
	convertData.AppendChar('{'); 
	convertData.NewLineWithTabs(); 
	
	convertData.IncrementNestCount()
	var forBlock *CodeBlock  = blockData.Block; ProcessBlock(convertData, forBlock, nestCount + 1, false); 
	convertData.DecrementNestCount()
	
	convertData.AppendChar('\r'); 
	convertData.AddTabs(); 
	convertData.AppendChar('}'); 
	convertData.NewLineWithTabs(); 
	
}
